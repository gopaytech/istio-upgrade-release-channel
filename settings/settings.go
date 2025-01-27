package settings

import (
	"github.com/kelseyhightower/envconfig"
)

type Settings struct {
	StorageMode               string `required:"true" envconfig:"STORAGE_MODE" default:"configmap"`
	StorageConfigMapName      string `envconfig:"STORAGE_CONFIGMAP_NAME" default:"istio-upgrade-release-channel"`
	StorageConfigMapNameSpace string `envconfig:"STORAGE_CONFIGMAP_NAMESPACE" default:"istio-system"`
}

func NewSettings() (Settings, error) {
	var settings Settings

	err := envconfig.Process("", &settings)
	if err != nil {
		return settings, err
	}

	return settings, nil
}
