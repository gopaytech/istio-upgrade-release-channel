package storage

import (
	"context"

	"github.com/gopaytech/istio-upgrade-release-channel/types"
)

type UpgradeInterface interface {
	Create(ctx context.Context, upgrade types.ReleaseChannel) error
	Get(ctx context.Context) (*types.ReleaseChannel, error)
	Update(ctx context.Context, upgrade types.ReleaseChannel) error
}
