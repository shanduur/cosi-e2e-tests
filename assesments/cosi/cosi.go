package cosi

import (
	"context"
	"testing"

	appsv1 "k8s.io/api/apps/v1"
	apiextensionsv1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	cosiv1alpha1 "sigs.k8s.io/container-object-storage-interface-api/apis/objectstorage/v1alpha1"
	"sigs.k8s.io/e2e-framework/klient/k8s/resources"
	"sigs.k8s.io/e2e-framework/pkg/envconf"
)

func CRDsInstalled(
	ctx context.Context,
	t *testing.T,
	cfg *envconf.Config,
) context.Context {
	var crds apiextensionsv1.CustomResourceDefinitionList

	expectedCRDs := []string{
		"bucketaccessclasses.objectstorage.k8s.io",
		"bucketaccesses.objectstorage.k8s.io",
		"bucketclaims.objectstorage.k8s.io",
		"bucketclasses.objectstorage.k8s.io",
		"buckets.objectstorage.k8s.io",
	}

	if err := cfg.Client().Resources().List(ctx, &crds); err != nil {
		t.Fatal(err)
	}

	found := 0

	for _, item := range crds.Items {
		for _, expected := range expectedCRDs {
			if item.GetObjectMeta().GetName() == expected {
				found++
			}
		}
	}

	if len(expectedCRDs) != found {
		t.Fatal("COSI CRDs not installed")
	}

	return ctx
}

func ObjectstorageControllerInstalled(
	ctx context.Context,
	t *testing.T,
	cfg *envconf.Config,
) context.Context {
	var deploymentList appsv1.DeploymentList

	selector := resources.WithLabelSelector("app.kubernetes.io/part-of=container-object-storage-interface")

	if err := cfg.Client().Resources().List(ctx, &deploymentList, selector); err != nil {
		t.Fatal(err)
	}

	if len(deploymentList.Items) == 0 {
		t.Fatal("deployment not found")
	}

	return ctx
}

func BucketAccessStatusReady(
	ctx context.Context,
	t *testing.T,
	cfg *envconf.Config,
) context.Context {
	return ctx
}

func BucketClaimStatusReady(
	ctx context.Context,
	t *testing.T,
	cfg *envconf.Config,
) context.Context {
	return ctx
}

func CreateBucket(bucket *cosiv1alpha1.Bucket) func(ctx context.Context, t *testing.T, cfg *envconf.Config) context.Context {
	return func(ctx context.Context, t *testing.T, cfg *envconf.Config) context.Context {
		return ctx
	}
}

func BucketExists(expected bool) func(ctx context.Context, t *testing.T, cfg *envconf.Config) context.Context {
	return func(ctx context.Context, t *testing.T, cfg *envconf.Config) context.Context {
		return ctx
	}
}

func CreateBucketClaim(bucketClaim *cosiv1alpha1.BucketClaim) func(ctx context.Context, t *testing.T, cfg *envconf.Config) context.Context {
	return func(ctx context.Context, t *testing.T, cfg *envconf.Config) context.Context {
		return ctx
	}
}

func DeleteBucketClaim(ctx context.Context, t *testing.T, cfg *envconf.Config) context.Context {
	return ctx
}

func CreateBucketAccess(bucketAccess *cosiv1alpha1.BucketAccess) func(ctx context.Context, t *testing.T, cfg *envconf.Config) context.Context {
	return func(ctx context.Context, t *testing.T, cfg *envconf.Config) context.Context {
		return ctx
	}
}

func SecretExists(expected bool) func(ctx context.Context, t *testing.T, cfg *envconf.Config) context.Context {
	return func(ctx context.Context, t *testing.T, cfg *envconf.Config) context.Context {
		return ctx
	}
}

func DeleteBucketAccess(ctx context.Context, t *testing.T, cfg *envconf.Config) context.Context {
	return ctx
}
