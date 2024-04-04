package cositest

import (
	"context"
	"e2e/assesments/cosi"
	"e2e/envfuncs"
	"flag"
	"log"
	"os"
	"testing"

	"sigs.k8s.io/container-object-storage-interface-api/apis/objectstorage/v1alpha1"
	"sigs.k8s.io/e2e-framework/pkg/env"
	"sigs.k8s.io/e2e-framework/pkg/envconf"
	"sigs.k8s.io/e2e-framework/pkg/features"
)

var (
	testenv env.Environment

	kind                  bool
	noInstallCRDs         bool
	noInstallController   bool
	noInstallSampleDriver bool
)

func TestMain(m *testing.M) {
	flag.BoolVar(&kind, "kind", false, "Start new environment with kind")
	flag.BoolVar(&noInstallCRDs, "cosi.no-install-crds", false, "Disable installing CRDs on cluster")
	flag.BoolVar(&noInstallController, "cosi.no-install-controller", false, "Disable installing COSI Controller on cluster")
	flag.BoolVar(&noInstallSampleDriver, "cosi.no-install-sample-driver", false, "Disable installing sample driver on cluster")

	cfg, err := envconf.NewFromFlags()
	if err != nil {
		log.Fatalf("failed to build envconf from flags: %s", err)
	}
	testenv = env.NewWithConfig(cfg)

	runID := envconf.RandomName("e2e-run", 32)

	testenv.Setup(
		// TODO: make this optional
		envfuncs.CreateCluster,
		envfuncs.InstallCRDs,
		envfuncs.InstallController,
		envfuncs.InstallDriver,
	)

	testenv.Finish(
		// TODO: make this optional
		envfuncs.UninstallDriver,
		envfuncs.UninstallController,
		envfuncs.UninstallCRDs,
		envfuncs.DeleteCluster,
	)

	testenv.BeforeEachTest(
		func(ctx context.Context, c *envconf.Config, t *testing.T) (context.Context, error) {
			return envfuncs.CreateNSForTest(ctx, c, t, runID)
		},
	)

	testenv.AfterEachTest(
		func(ctx context.Context, c *envconf.Config, t *testing.T) (context.Context, error) {
			return envfuncs.DeleteNSForTest(ctx, c, t, runID)
		},
	)

	os.Exit(testenv.Run(m))
}

func TestBucketProvisioning(t *testing.T) {
	testenv.Test(t,
		features.New("Greenfield Bucket").
			Assess("BucketClass is created",
				cosi.CreateBucketClass(&v1alpha1.BucketClass{})).
			Assess("BucketClaim is created",
				cosi.CreateBucketClaim(&v1alpha1.BucketClaim{})).
			Assess("Bucket is created",
				cosi.CreateBucket(&v1alpha1.Bucket{})).
			Assess("BucketClaim has ready status",
				cosi.BucketClaimStatusReady).
			Assess("BucketClaim is deleted",
				cosi.DeleteBucketClaim).
			Assess("Bucket is deleted",
				cosi.BucketExists(false)).
			Feature(),

		features.New("Brownfield Bucket").
			Assess("BucketClass is created",
				cosi.CreateBucketClass(&v1alpha1.BucketClass{})).
			Assess("BucketClaim is created",
				cosi.CreateBucketClaim(&v1alpha1.BucketClaim{})).
			Assess("Bucket is created",
				cosi.CreateBucket(&v1alpha1.Bucket{})).
			Assess("BucketClaim has ready status",
				cosi.BucketClaimStatusReady).
			Assess("BucketClaim is deleted",
				cosi.DeleteBucketClaim).
			Assess("Bucket is deleted",
				cosi.BucketExists(false)).
			Feature(),
	)
}

func TestBucketAccessProvisioning(t *testing.T) {
	testenv.Test(t,
		features.New("BucketAccess").
			Assess("BucketClass is created",
				cosi.CreateBucketClass(&v1alpha1.BucketClass{})).
			Assess("BucketClaim is created",
				cosi.CreateBucketClaim(&v1alpha1.BucketClaim{})).
			Assess("Bucket is created",
				cosi.CreateBucket(&v1alpha1.Bucket{})).
			Assess("BucketClaim has ready status",
				cosi.BucketClaimStatusReady).
			Assess("BucketAccessClass is created",
				cosi.CreateBucketAccessClass(&v1alpha1.BucketAccessClass{})).
			Assess("BucketAccess is created",
				cosi.CreateBucketAccess(&v1alpha1.BucketAccess{})).
			Assess("BucketAccess has ready status",
				cosi.BucketAccessStatusReady).
			Assess("BucketAccess is deleted",
				cosi.DeleteBucketAccess).
			Assess("Secret is deleted",
				cosi.SecretExists(false)).
			Assess("BucketClaim is deleted",
				cosi.DeleteBucketClaim).
			Assess("Bucket is deleted",
				cosi.BucketExists(false)).
			Feature(),
	)
}
