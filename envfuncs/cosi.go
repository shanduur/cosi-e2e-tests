package envfuncs

import (
	"context"

	"sigs.k8s.io/e2e-framework/pkg/envconf"
)

func InstallCRDs(ctx context.Context, cfg *envconf.Config) (context.Context, error) {
	return ctx, nil
}

func UninstallCRDs(ctx context.Context, cfg *envconf.Config) (context.Context, error) {
	return ctx, nil
}

func InstallController(ctx context.Context, cfg *envconf.Config) (context.Context, error) {
	return ctx, nil
}

func UninstallController(ctx context.Context, cfg *envconf.Config) (context.Context, error) {
	return ctx, nil
}

func InstallDriver(ctx context.Context, cfg *envconf.Config) (context.Context, error) {
	return ctx, nil
}

func UninstallDriver(ctx context.Context, cfg *envconf.Config) (context.Context, error) {
	return ctx, nil
}

func CreateBucketClass(ctx context.Context, cfg *envconf.Config) (context.Context, error) {
	return ctx, nil
}

func DeleteBucketClass(ctx context.Context, cfg *envconf.Config) (context.Context, error) {
	return ctx, nil
}

func CreateBucketAccessClass(ctx context.Context, cfg *envconf.Config) (context.Context, error) {
	return ctx, nil
}

func DeleteBucketAccessClass(ctx context.Context, cfg *envconf.Config) (context.Context, error) {
	return ctx, nil
}
