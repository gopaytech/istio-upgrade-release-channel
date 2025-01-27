package main

import (
	"log"

	"github.com/gopaytech/istio-upgrade-release-channel/services/storage"
	"github.com/gopaytech/istio-upgrade-release-channel/settings"
)

func main() {
	settings, err := settings.NewSettings()
	if err != nil {
		log.Fatal(err)
	}

	_, err = storage.ReleaseChannelFactory(settings)
	if err != nil {
		log.Fatal(err)
	}
}
