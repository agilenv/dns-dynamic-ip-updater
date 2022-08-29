package main

import (
	"github.com/agilenv/linkip/internal/dns"
	"github.com/agilenv/linkip/internal/dns/provider"
	"github.com/agilenv/linkip/internal/dns/publicip"
	"github.com/agilenv/linkip/internal/dns/track"
	"github.com/agilenv/linkip/pkg/rest"
)

func buildUpdater() *dns.Updater {
	trackFileStorage := track.NewFileStorage("tracks.log")
	IpifyAPI := publicip.NewIpifyPublicIPAPI(rest.NewClient())
	digitaloceanDNSProvider, err := provider.NewDigitaloceanProvider(rest.NewClient())
	if err != nil {
		panic(err)
	}
	return dns.NewUpdater(digitaloceanDNSProvider, trackFileStorage, IpifyAPI)
}
