package storage

import (
	"fmt"

	"github.com/gopaytech/istio-upgrade-release-channel/settings"
	"github.com/gopaytech/istio-upgrade-release-channel/storage/postgresql"
)

func ReleaseChannelFactory(settings settings.Settings) (ReleaseChannelInterface, error) {
	if settings.StorageMode == "postgresql" {
		return postgresql.NewPostgresqlStorage(settings)
	}

	if settings.StorageMode == "configmap" {
		return nil, fmt.Errorf("storage is not supported %s", settings.StorageMode)
	}

	return nil, fmt.Errorf("storage is not supported %s", settings.StorageMode)
}
