package models

import (
	"time"

	"github.com/hashicorp/go-version"
)

type ReleaseChannel struct {
	ImmediateVersions []Version `json:"immediate_versions" yaml:"immediate_versions"`
	RapidVersion      []Version `json:"rapid_version" yaml:"rapid_version"`
	RegularVersion    []Version `json:"regular_version" yaml:"regular_version"`
	StableVersion     []Version `json:"stable_version" yaml:"stable_version"`
}

type Version struct {
	Version       version.Version
	PromotionDate time.Time
}
