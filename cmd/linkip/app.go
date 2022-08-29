package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/urfave/cli/v2"
)

const (
	envFilepath = ".env"
)

func buildApp() *cli.App {
	envFile := envFilepath
	return &cli.App{
		Name:  "linkip",
		Usage: "Simple application to keep updated a dns record where the ip associated is dynamic",
		Commands: []*cli.Command{
			statusCMD(),
			updateCMD(),
			listCMD(),
		},
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "env-file",
				Usage:       "path to env file",
				Destination: &envFile,
				Value:       envFile,
				HasBeenSet:  true,
			},
		},
		Before: func(cCtx *cli.Context) error {
			if err := godotenv.Load(envFile); err != nil {
				fmt.Fprintf(os.Stderr, "%s\n", err)
				return errors.New("missing .env file. You can provide a file path with the flag --env-file")
			}
			return nil
		},
	}
}
