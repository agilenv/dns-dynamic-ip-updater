package main

import (
	"fmt"
	"os"

	"github.com/agilenv/linkip/internal/dns"
	"github.com/agilenv/linkip/internal/dns/provider"
	"github.com/agilenv/linkip/internal/dns/publicip"
	"github.com/agilenv/linkip/internal/dns/track"
	"github.com/agilenv/linkip/pkg/rest"
)

const (
	digitalOceanProvider = "digitalocean"
)

var (
	availableDNSProviders = []string{
		digitalOceanProvider,
	}
)

func buildFileStats() *dns.Stats {
	trackFileStorage := track.NewFileStorage()
	return dns.NewStats(trackFileStorage)
}

func buildUpdater(dnsProvider string) *dns.Updater {
	var u *dns.Updater
	fileStats := buildFileStats()
	IpifyAPI := publicip.NewIpifyPublicIPAPI(rest.NewClient())
	switch dnsProvider {
	case digitalOceanProvider:
		p, err := provider.NewDigitaloceanProvider(rest.NewClient())
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s", err)
			os.Exit(1)
		}
		u = dns.NewUpdater(p, fileStats, IpifyAPI)
	default:
		fmt.Fprintf(os.Stderr, "invalid dns provider")
		os.Exit(1)
	}
	return u
}
