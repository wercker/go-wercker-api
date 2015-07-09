package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/codegangsta/cli"
	"github.com/wercker/go-wercker-api"
	"github.com/wercker/go-wercker-api/credentials"
)

var (
	buildsCommand = cli.Command{
		Name:  "builds",
		Usage: "build related endpoints",
		Subcommands: []cli.Command{
			cli.Command{
				Name:  "fetch",
				Usage: "retrieve a single build",
				Action: wrapper(func(c *cli.Context, client *wercker.Client) (interface{}, error) {
					buildID := c.Args().First()
					if buildID == "" {
						return nil, fmt.Errorf("build id is required as an argument")
					}

					getBuildOptions := &wercker.GetBuildOptions{
						BuildID: buildID,
					}

					return client.Builds.Get(getBuildOptions)
				}),
			},
		},
	}
	applicationsCommand = cli.Command{
		Name:  "applications",
		Usage: "application related endpoints",
		Subcommands: []cli.Command{
			cli.Command{
				Name:  "fetch",
				Usage: "retrieve a single application",
				Action: wrapper(func(c *cli.Context, client *wercker.Client) (interface{}, error) {
					owner := c.Args().First()
					name := c.Args().Get(1)

					if owner == "" {
						return nil, fmt.Errorf("owner is required as the first argument")
					}

					if name == "" {
						s := strings.SplitN(owner, "/", 2)
						if len(s) != 2 {
							return nil, fmt.Errorf("application name is required as the second argument")
						}
						owner = s[0]
						name = s[1]
					}

					getApplicationOptions := &wercker.GetApplicationOptions{
						Owner: owner,
						Name:  name,
					}

					return client.Applications.Get(getApplicationOptions)
				}),
			},
			cli.Command{
				Name:  "builds",
				Usage: "retrieves the builds for an application",
				Flags: []cli.Flag{
					cli.StringFlag{Name: "branch"},
					cli.StringFlag{Name: "commit"},
					cli.StringFlag{Name: "limit"},
					cli.StringFlag{Name: "result"},
					cli.StringFlag{Name: "skip"},
					cli.StringFlag{Name: "sort"},
					cli.StringFlag{Name: "stack"},
					cli.StringFlag{Name: "status"},
				},
				Action: wrapper(func(c *cli.Context, client *wercker.Client) (interface{}, error) {
					owner := c.Args().First()
					name := c.Args().Get(1)

					if owner == "" {
						return nil, fmt.Errorf("owner is required as the first argument")
					}

					if name == "" {
						s := strings.SplitN(owner, "/", 2)
						if len(s) != 2 {
							return nil, fmt.Errorf("application name is required as the second argument")
						}
						owner = s[0]
						name = s[1]
					}

					fetchForApplicationOptions := &wercker.FetchForApplicationOptions{
						Owner:  owner,
						Name:   name,
						Branch: c.String("branch"),
						Commit: c.String("commit"),
						Limit:  c.String("limit"),
						Result: c.String("result"),
						Skip:   c.String("skip"),
						Sort:   c.String("sort"),
						Stack:  c.String("stack"),
						Status: c.String("status"),
					}

					return client.Builds.FetchForApplication(fetchForApplicationOptions)
				}),
			},
		},
	}
)

func wrapper(f func(c *cli.Context, client *wercker.Client) (interface{}, error)) func(c *cli.Context) {
	return func(c *cli.Context) {
		client := createClient(c)

		result, err := f(c, client)
		if err != nil {
			os.Stderr.WriteString("Unable to fetch from the API: ")
			os.Stderr.WriteString(err.Error())
			os.Stderr.WriteString("\n")
			return
		}

		b, err := json.MarshalIndent(result, "", "  ")
		if err != nil {
			os.Stderr.WriteString("Unable to marshal response from the API: ")
			os.Stderr.WriteString(err.Error())
			os.Stderr.WriteString("\n")
			return
		}

		os.Stdout.Write(b)
		os.Stdout.WriteString("\n")
	}
}

func createClient(c *cli.Context) *wercker.Client {
	endpoint := c.GlobalString("endpoint")
	config := &wercker.Config{
		Endpoint: endpoint,
	}

	if c.GlobalBool("anonymous") {
		config.Credentials = credentials.Anonymous()
	} else {
		token := c.GlobalString("token")
		if token != "" {
			config.Credentials = credentials.Token(token)
		}
	}

	client := wercker.NewClient(config)

	return client
}

func main() {
	app := cli.NewApp()

	app.Author = "wercker"
	app.Email = "pleasemailus@wercker.com"
	app.Name = "api explorer"
	app.Usage = "retrieve results from the wercker API"
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
			Value: "",
			Usage: "Token used for authentication (leave empty to use default SDK strategy)",
		},
		cli.BoolFlag{
			Name:  "anonymous",
			Usage: "Force the call to use anonymous credentials",
		},
	}
	app.Commands = []cli.Command{
		applicationsCommand,
		buildsCommand,
	}

	app.Run(os.Args)
}
