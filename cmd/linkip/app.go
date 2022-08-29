package main

import (
	"github.com/urfave/cli/v2"
)

func buildApp() *cli.App {
	return &cli.App{
		Name:  "linkip",
		Usage: "Simple application to keep updated a dns record where the ip associated is dynamic",
		Commands: []*cli.Command{
			statusCMD(),
			updateCMD(),
			listCMD(),
		},
	}
}
