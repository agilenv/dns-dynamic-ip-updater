package main

import (
	"log"
	"os"

	"github.com/agilenv/linkip/internal/dns"
	"github.com/agilenv/linkip/internal/dns/provider"
	"github.com/agilenv/linkip/internal/dns/publicip"
	"github.com/agilenv/linkip/internal/dns/track"
	"github.com/agilenv/linkip/pkg/rest"
)

func buildStats() *dns.Stats {
	trackFileStorage := track.NewFileStorage()
	return dns.NewStats(trackFileStorage)
}

func buildUpdater() *dns.Updater {
	stats := buildStats()
	IpifyAPI := publicip.NewIpifyPublicIPAPI(rest.NewClient())
	digitaloceanDNSProvider, err := provider.NewDigitaloceanProvider(rest.NewClient())
	if err != nil {
		log.Fatalf("%s", err)
		os.Exit(1)
	}
	return dns.NewUpdater(digitaloceanDNSProvider, stats, IpifyAPI)
}
