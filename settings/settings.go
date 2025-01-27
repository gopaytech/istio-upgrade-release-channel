package settings

import (
	"errors"

	"github.com/kelseyhightower/envconfig"
)

var (
	ErrStoragePostgresqlHostEmpty     = errors.New("storage postgresql host is empty")
	ErrStoragePostgresqlPortEmpty     = errors.New("storage postgresql port is empty")
	ErrStoragePostgresqlUserEmpty     = errors.New("storage postgresql user is empty")
	ErrStoragePostgresqlPasswordEmpty = errors.New("storage postgresql password is empty")
	ErrStoragePostgresqlDatabaseEmpty = errors.New("storage postgresql database is empty")
)

type Settings struct {
	StorageMode               string `required:"true" envconfig:"STORAGE_MODE" default:"postgresql"`
	StoragePostgresqlHost     string `envconfig:"STORAGE_POSTGRESQL_HOST"`
	StoragePostgresqlPort     string `envconfig:"STORAGE_POSTGRESQL_PORT"`
	StoragePostgresqlUser     string `envconfig:"STORAGE_POSTGRESQL_USER"`
	StoragePostgresqlPassword string `envconfig:"STORAGE_POSTGRESQL_PASSWORD"`
	StoragePostgresqlDatabase string `envconfig:"STORAGE_POSTGRESQL_DATABASE"`
	StorageConfigMapName      string `envconfig:"STORAGE_CONFIGMAP_NAME" default:"istio-upgrade"`
	StorageConfigMapNameSpace string `envconfig:"STORAGE_CONFIGMAP_NAMESPACE" default:"istio-system"`
}

func (s Settings) Validate() error {
	if s.StorageMode == "postgresql" {
		if s.StoragePostgresqlHost == "" {
			return ErrStoragePostgresqlHostEmpty
		}
		if s.StoragePostgresqlPort == "" {
			return ErrStoragePostgresqlPortEmpty
		}
		if s.StoragePostgresqlUser == "" {
			return ErrStoragePostgresqlUserEmpty
		}
		if s.StoragePostgresqlPassword == "" {
			return ErrStoragePostgresqlPasswordEmpty
		}
		if s.StoragePostgresqlDatabase == "" {
			return ErrStoragePostgresqlDatabaseEmpty
		}
	}

	return nil
}

func NewSettings() (Settings, error) {
	var settings Settings

	err := envconfig.Process("", &settings)
	if err != nil {
		return settings, err
	}

	return settings, nil
}
