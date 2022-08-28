package main

import (
	"context"
	"fmt"

	"github.com/agilenv/dns-dynamic-ip-updater/internal/dns"
	"github.com/agilenv/dns-dynamic-ip-updater/internal/dns/ip"
	"github.com/agilenv/dns-dynamic-ip-updater/internal/dns/track"

	"github.com/agilenv/dns-dynamic-ip-updater/internal/dns/provider"
	"github.com/agilenv/dns-dynamic-ip-updater/pkg/rest"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	// TRACK
	t := track.NewFileStorage("tracks.log")

	// PUBLIC IP
	rPublicIP := rest.NewClient()
	pIP := ip.NewIpifyPublicIPAPI(rPublicIP)

	// PROVIDER
	rProvider := rest.NewClient()
	do, err := provider.NewDigitaloceanProvider(rProvider)
	if err != nil {
		panic(err)
	}

	u := dns.NewUpdater(do, t, pIP)
	fmt.Println(u.Sync(context.Background()))
}
