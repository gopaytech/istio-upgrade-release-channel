package storage

import (
	"context"

	"github.com/gopaytech/istio-upgrade-release-channel/models"
)

type ReleaseChannelInterface interface {
	Create(ctx context.Context, releaseChannel models.ReleaseChannel) error
	Get(ctx context.Context) (*models.ReleaseChannel, error)
	Update(ctx context.Context, releaseChannel models.ReleaseChannel) error
}
