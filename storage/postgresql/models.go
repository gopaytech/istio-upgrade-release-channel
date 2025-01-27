package postgresql

import (
	"time"

	"github.com/google/uuid"
	"github.com/hashicorp/go-version"
	"gorm.io/gorm"
)

type ReleaseChannel struct {
	ID                uuid.UUID  `gorm:"type:uuid;not null" json:"id"`
	CreatedAt         time.Time  `json:"-"`
	UpdatedAt         *time.Time `json:"-"`
	DeletedAt         *time.Time `gorm:"index" json:"-"`
	ImmediateVersions []Version  `gorm:"type:jsonb" json:"immediate_versions"`
	RapidVersion      []Version  `gorm:"type:jsonb" json:"rapid_version"`
	RegularVersion    []Version  `gorm:"type:jsonb" json:"regular_version"`
	StableVersion     []Version  `gorm:"type:jsonb" json:"stable_version"`
}

func (r *ReleaseChannel) BeforeCreate(tx *gorm.DB) error {
	r.ID = uuid.New()
	return nil
}

func (r *ReleaseChannel) BeforeSave(tx *gorm.DB) (err error) {
	updatedAt := time.Now()
	r.UpdatedAt = &updatedAt
	return nil
}

type Version struct {
	Version       version.Version
	PromotionDate time.Time
}
