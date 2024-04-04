package envfuncs

import (
	"context"
	"e2e/retry"
	"fmt"

	"k8s.io/client-go/kubernetes"
	"sigs.k8s.io/e2e-framework/pkg/envconf"
	"sigs.k8s.io/e2e-framework/support/kind"
)

type (
	ClusterCtxKey string
)

func CreateCluster(ctx context.Context, cfg *envconf.Config) (context.Context, error) {
	name := envconf.RandomName("cosi-e2e-cluster", 32)
	cluster := kind.NewCluster(name)
	kubeconfig, err := cluster.Create(ctx)
	if err != nil {
		return ctx, err
	}

	// stall a bit to allow most pods to come up
	ctx, err = IsClusterReady(ctx, cfg)
	if err != nil {
		return ctx, err
	}

	// update environment with kubecofig file
	cfg.WithKubeconfigFile(kubeconfig)

	// propagate cluster value
	return context.WithValue(ctx, ClusterCtxKey("cluster"), cluster), nil
}

func IsClusterReady(ctx context.Context, cfg *envconf.Config) (context.Context, error) {
	clientset, err := kubernetes.NewForConfig(cfg.Client().RESTConfig())
	if err != nil {
		return ctx, fmt.Errorf("failed to create clientset from klient: %w", err)
	}

	return ctx, retry.Do(func() error {
		_, err := clientset.ServerVersion()
		return err
	})
}

func DeleteCluster(ctx context.Context, _ *envconf.Config) (context.Context, error) {
	cluster := ctx.Value(ClusterCtxKey("cluster")).(*kind.Cluster)
	if cluster == nil {
		return ctx, fmt.Errorf("error getting kind cluster from context")
	}

	if err := cluster.Destroy(ctx); err != nil {
		return ctx, err
	}

	return ctx, nil
}
