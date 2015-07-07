package main

import (
	"fmt"
	"os"

	"github.com/codegangsta/cli"
	"github.com/wercker/go-wercker-api"
	"github.com/wercker/go-wercker-api/credentials"
)

var (
	applicationsCommand = cli.Command{
		Name:  "builds",
		Usage: "builds related endpoints",

		Subcommands: []cli.Command{
			cli.Command{
				Name:  "fetch",
				Usage: "retrieve a single build",
				Action: func(c *cli.Context) {

					endpoint := c.GlobalString("endpoint")
					options := wercker.Options{
						Endpoint: endpoint,
					}

					token := c.GlobalString("token")
					if token != "" {
						options.Creds = credentials.Token(token)
					}
					client := wercker.NewClient(&options)

					getBuildOptions := &wercker.GetBuildOptions{
						BuildID: "123",
					}

					result, err := client.GetBuild(getBuildOptions)
					if err != nil {
						panic(err)
					}
					fmt.Printf("result: %v", result)
				},
			},
		},
	}
)

func main() {
	app := cli.NewApp()

	app.Author = "wercker"
	app.Email = "pleasemailus@wercker.com"
	app.Name = "api-client"
	app.Usage = "A new cli application"
	app.Version = FullVersion()

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:   "endpoint",
			Value:  "https://app.wercker.com",
			Usage:  "Base url for the wercker app.",
			EnvVar: "WERCKER_ENDPOINT",
		},
		cli.StringFlag{
			Name:  "token",
			Value: "https://app.wercker.com",
			Usage: "Base url for the wercker app.",
		},
	}
	app.Commands = []cli.Command{
		applicationsCommand,
	}

	app.Run(os.Args)
}
