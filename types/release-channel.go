package types

import "github.com/hashicorp/go-version"

type ReleaseChannel struct {
	ImmediateVersion version.Version
	RapidVersion     version.Version
	RegularVersion   version.Version
	StableVersion    version.Version
}

func (r ReleaseChannel) ToConfigMapData() map[string]string {
	return map[string]string{
		"immediate_version": r.ImmediateVersion.String(),
		"rapid_version":     r.RapidVersion.String(),
		"regular_version":   r.RegularVersion.String(),
		"stable_version":    r.StableVersion.String(),
	}
}
