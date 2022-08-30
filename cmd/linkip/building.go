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
	filepath := os.Getenv(trackFilepath)
	if filepath == "" {
		filepath = "linkip_tracks.log"
	}
	trackFileStorage := track.NewFileStorage(filepath)
	return dns.NewStats(trackFileStorage)
}

func buildUpdater(dnsProvider string) *dns.Updater {
	var u *dns.Updater
	fileStats := buildFileStats()
	IpifyAPI := publicip.NewIpifyPublicIPAPI(rest.NewClient())
	switch dnsProvider {
	case digitalOceanProvider:
		p := provider.NewDigitaloceanProvider(rest.NewClient(), provider.DigitaloceanConfig{
			DomainName: os.Getenv(digitalOceanDomainName),
			RecordID:   os.Getenv(digitalOceanRecordID),
			Token:      os.Getenv(digitalOceanToken),
		})
		u = dns.NewUpdater(p, fileStats, IpifyAPI)
	default:
		fmt.Fprintf(os.Stderr, "invalid dns provider")
		os.Exit(1)
	}
	return u
}
