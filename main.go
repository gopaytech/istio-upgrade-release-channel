package main

import (
	"log"

	"github.com/gopaytech/istio-upgrade-release-channel/settings"
	"github.com/gopaytech/istio-upgrade-release-channel/storage"
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
