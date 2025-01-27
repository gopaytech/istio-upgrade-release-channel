package postgresql

import (
	"context"
	"fmt"

	"github.com/gopaytech/istio-upgrade-release-channel/models"
	"github.com/gopaytech/istio-upgrade-release-channel/settings"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type PostgresqlStorage struct {
	db *gorm.DB
}

func (p PostgresqlStorage) Create(ctx context.Context, releaseChannel models.ReleaseChannel) error {
	rc := toPostgresqlReleaseChannel(releaseChannel)
	err := p.db.Create(&rc).Error

	if err != nil {
		return err
	}
	return nil
}

func (p PostgresqlStorage) Get(ctx context.Context) (*models.ReleaseChannel, error) {
	rc, err := p.getPostgresqlReleaseChannel(ctx)
	if err != nil {
		return nil, err
	}

	modelRC := toReleaseChannel(*rc)
	return &modelRC, nil
}

func (p PostgresqlStorage) getPostgresqlReleaseChannel(ctx context.Context) (*ReleaseChannel, error) {
	var rc ReleaseChannel

	err := p.db.First(&rc).Error
	if err != nil {
		return nil, err
	}

	return &rc, nil
}

func (p PostgresqlStorage) Update(ctx context.Context, releaseChannel models.ReleaseChannel) error {
	updatedRC := toPostgresqlReleaseChannel(releaseChannel)
	existingRC, err := p.getPostgresqlReleaseChannel(ctx)
	if err != nil {
		return err
	}

	existingRC.ImmediateVersions = updatedRC.ImmediateVersions
	existingRC.RapidVersion = updatedRC.RapidVersion
	existingRC.RegularVersion = updatedRC.RegularVersion
	existingRC.StableVersion = updatedRC.StableVersion

	err = p.db.Save(&existingRC).Error
	if err != nil {
		return err
	}

	return nil
}

const (
	defaultTimeZone = "Asia/Jakarta"
)

const dsnTemplate = "host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=%s"

func NewPostgresqlStorage(settings settings.Settings) (PostgresqlStorage, error) {
	dsn := fmt.Sprintf(dsnTemplate, settings.StoragePostgresqlHost, settings.StoragePostgresqlUser, settings.StoragePostgresqlPassword, settings.StoragePostgresqlDatabase, settings.StoragePostgresqlPort, defaultTimeZone)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return PostgresqlStorage{}, err
	}

	err = db.AutoMigrate(&models.ReleaseChannel{})
	if err != nil {
		return PostgresqlStorage{}, err
	}

	return PostgresqlStorage{
		db: db,
	}, nil
}

func toPostgresqlReleaseChannel(releaseChannel models.ReleaseChannel) ReleaseChannel {
	rc := ReleaseChannel{}

	for _, v := range releaseChannel.ImmediateVersions {
		rc.ImmediateVersions = append(rc.ImmediateVersions, Version{
			Version:       v.Version,
			PromotionDate: v.PromotionDate,
		})
	}

	for _, v := range releaseChannel.RapidVersion {
		rc.RapidVersion = append(rc.RapidVersion, Version{
			Version:       v.Version,
			PromotionDate: v.PromotionDate,
		})
	}

	for _, v := range releaseChannel.RegularVersion {
		rc.RegularVersion = append(rc.RegularVersion, Version{
			Version:       v.Version,
			PromotionDate: v.PromotionDate,
		})
	}

	for _, v := range releaseChannel.StableVersion {
		rc.StableVersion = append(rc.StableVersion, Version{
			Version:       v.Version,
			PromotionDate: v.PromotionDate,
		})
	}

	return rc
}

func toReleaseChannel(releaseChannel ReleaseChannel) models.ReleaseChannel {
	rc := models.ReleaseChannel{}

	for _, v := range releaseChannel.ImmediateVersions {
		rc.ImmediateVersions = append(rc.ImmediateVersions, models.Version{
			Version:       v.Version,
			PromotionDate: v.PromotionDate,
		})
	}

	for _, v := range releaseChannel.RapidVersion {
		rc.RapidVersion = append(rc.RapidVersion, models.Version{
			Version:       v.Version,
			PromotionDate: v.PromotionDate,
		})
	}

	for _, v := range releaseChannel.RegularVersion {
		rc.RegularVersion = append(rc.RegularVersion, models.Version{
			Version:       v.Version,
			PromotionDate: v.PromotionDate,
		})
	}

	for _, v := range releaseChannel.StableVersion {
		rc.StableVersion = append(rc.StableVersion, models.Version{
			Version:       v.Version,
			PromotionDate: v.PromotionDate,
		})
	}

	return rc
}
