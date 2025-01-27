package storage

import (
	"fmt"

	"github.com/gopaytech/istio-upgrade-release-channel/config"
	"github.com/gopaytech/istio-upgrade-release-channel/services/storage/configmap"
	"github.com/gopaytech/istio-upgrade-release-channel/settings"
)

func ReleaseChannelFactory(settings settings.Settings) (UpgradeInterface, error) {
	if settings.StorageMode == "configmap" {
		kubernetesConfig, err := config.LoadKubernetes()
		if err != nil {
			return nil, err
		}

		return configmap.NewReleaseChannelConfigMap(kubernetesConfig, settings), nil
	}

	return nil, fmt.Errorf("storage is not supported %s", settings.StorageMode)
}
