package main

import (
	"github.com/urfave/cli/v2"
)

func buildApp() *cli.App {
	u := buildUpdater()

	return &cli.App{
		Name:  "linkip",
		Usage: "Application that helps to synchronize a dynamic IP to a DNS record",
		Commands: []*cli.Command{
			statusCMD(u),
			updateCMD(u),
		},
	}
}
