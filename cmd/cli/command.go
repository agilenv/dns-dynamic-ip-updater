package main

import (
	"context"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/agilenv/linkip/internal/dns/track"

	"github.com/urfave/cli/v2"
)

type updater interface {
	SearchForChanges(ctx context.Context) (bool, string, error)
	Update(ctx context.Context, ip string) error
	LastExecution() *track.Event
}

func updateCMD(u updater) *cli.Command {
	var confirm string
	return &cli.Command{
		Name:  "update",
		Usage: "Search for IP changes and update a DNS record",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "confirm",
				Usage: "updates to dns record on provider [yes/no]",
				Value: "",

				Destination: &confirm,
			},
		},
		Action: func(cCtx *cli.Context) error {
			ctx := context.Background()
			changed, ip, err := u.SearchForChanges(ctx)
			if err != nil {
				return errors.New("an unexpected error occurred when fetching public IP address")
			}
			if changed {
				fmt.Fprintf(os.Stdout, "New public IP [%s] has founded\n", ip)
				if confirm == "" {
					fmt.Fprintf(os.Stdout, "Do you want to update the dns record? [yes]: ")
					confirm = "yes"
					fmt.Scanf("%s", &confirm)
				}
				if confirm == "yes" {
					if err = u.Update(ctx, ip); err != nil {
						return errors.New("cannot update DNS Record. [" + err.Error() + "]")
					}
					fmt.Fprintf(os.Stdout, "Done!\n")
					return nil
				}
				return nil
			}
			fmt.Fprintf(os.Stdout, "No changes from last execution, IP [%s]\n", ip)
			return nil
		},
	}
}

func statusCMD(u updater) *cli.Command {
	return &cli.Command{
		Name:  "status",
		Usage: "Get information from last execution",
		Action: func(ctx *cli.Context) error {
			if event := u.LastExecution(); event != nil {
				fmt.Fprintf(os.Stdout, "\nLast Execution:\n\t%s\n", event.Time.Format(time.RFC850))
				fmt.Fprintf(os.Stdout, "\nIP Address:\n\t%s\n", event.IP)
				fmt.Fprintf(os.Stdout, "\nPublic IP API:\n\t%s\n", event.PublicAPI)
				return nil
			}
			fmt.Fprintln(os.Stdout, "No records found.")
			return nil
		},
	}
}
