package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/codegangsta/cli"
	"github.com/wercker/go-wercker-api"
	"github.com/wercker/go-wercker-api/credentials"
)

var (
	tokensCommand = cli.Command{
		Name:  "tokens",
		Usage: "tokens related endpoints",
		Subcommands: []cli.Command{
			cli.Command{
				Name:  "list",
				Usage: "retrieve all tokens for the current user",
				Action: wrapper(func(c *cli.Context, client *wercker.Client) (interface{}, error) {
					getTokensOptions := &wercker.GetTokensOptions{}
					return client.GetTokens(getTokensOptions)
				}),
			},
			cli.Command{
				Name:  "get",
				Usage: "retrieve a single token",
				Action: wrapper(func(c *cli.Context, client *wercker.Client) (interface{}, error) {
					tokenID := c.Args().First()
					if tokenID == "" {
						return nil, fmt.Errorf("token id is required as an argument")
					}
					getTokenOptions := &wercker.GetTokenOptions{TokenID: tokenID}
					return client.GetToken(getTokenOptions)
				}),
			},
			cli.Command{
				Name:  "delete",
				Usage: "delete a single token",
				Action: wrapper(func(c *cli.Context, client *wercker.Client) (interface{}, error) {
					tokenID := c.Args().First()
					if tokenID == "" {
						return nil, fmt.Errorf("token id is required as an argument")
					}
					deleteTokenOptions := &wercker.DeleteTokenOptions{TokenID: tokenID}
					return nil, client.DeleteToken(deleteTokenOptions)
				}),
			},
			cli.Command{
				Name:  "create",
				Usage: "create a single token",
				Flags: []cli.Flag{
					cli.StringFlag{Name: "name"},
				},
				Action: wrapper(func(c *cli.Context, client *wercker.Client) (interface{}, error) {
					name := c.String("name")
					if name == "" {
						return nil, errors.New("name is required")
					}
					createTokenOptions := &wercker.CreateTokenOptions{Name: name}
					return client.CreateToken(createTokenOptions)
				}),
			},
			cli.Command{
				Name:  "update",
				Usage: "update a single token",
				Flags: []cli.Flag{
					cli.StringFlag{Name: "name"},
				},
				Action: wrapper(func(c *cli.Context, client *wercker.Client) (interface{}, error) {
					tokenID := c.Args().First()
					if tokenID == "" {
						return nil, fmt.Errorf("token id is required as an argument")
					}
					name := c.String("name")
					updateTokenOptions := &wercker.UpdateTokenOptions{TokenID: tokenID, Name: name}
					return client.UpdateToken(updateTokenOptions)
				}),
			},
		},
	}

	buildCommand = cli.Command{
		Name:  "builds",
		Usage: "build related endpoints",
		Subcommands: []cli.Command{
			cli.Command{
				Name:  "get",
				Usage: "retrieve a single build",
				Action: wrapper(func(c *cli.Context, client *wercker.Client) (interface{}, error) {
					buildID := c.Args().First()
					if buildID == "" {
						return nil, fmt.Errorf("build id is required as an argument")
					}
					getBuildOptions := &wercker.GetBuildOptions{BuildID: buildID}
					return client.GetBuild(getBuildOptions)
				}),
			},
			cli.Command{
				Name:  "list",
				Usage: "retrieves the builds for an application",
				Flags: []cli.Flag{
					cli.StringFlag{Name: "branch"},
					cli.StringFlag{Name: "commit"},
					cli.IntFlag{Name: "limit"},
					cli.StringFlag{Name: "result"},
					cli.IntFlag{Name: "skip"},
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

					getBuildsOptions := &wercker.GetBuildsOptions{
						Owner:  owner,
						Name:   name,
						Branch: c.String("branch"),
						Commit: c.String("commit"),
						Limit:  c.Int("limit"),
						Result: c.String("result"),
						Skip:   c.Int("skip"),
						Sort:   c.String("sort"),
						Stack:  c.String("stack"),
						Status: c.String("status"),
					}

					return client.GetBuilds(getBuildsOptions)
				}),
			},
		},
	}
	deployCommand = cli.Command{
		Name:  "deploy",
		Usage: "deploy related endpoints",
		Subcommands: []cli.Command{
			cli.Command{
				Name:  "get",
				Usage: "retrieve a single deploy",
				Action: wrapper(func(c *cli.Context, client *wercker.Client) (interface{}, error) {
					deployID := c.Args().First()
					if deployID == "" {
						return nil, fmt.Errorf("deploy id is required as an argument")
					}
					getDeployOptions := &wercker.GetDeployOptions{DeployID: deployID}
					return client.GetDeploy(getDeployOptions)
				}),
			},
			cli.Command{
				Name:  "list",
				Usage: "retrieves the deploys for an application",
				Flags: []cli.Flag{
					cli.StringFlag{Name: "build-id"},
					cli.IntFlag{Name: "limit"},
					cli.StringFlag{Name: "result"},
					cli.IntFlag{Name: "skip"},
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

					getDeploysOptions := &wercker.GetDeploysOptions{
						Owner:   owner,
						Name:    name,
						BuildID: c.String("build-id"),
						Limit:   c.Int("limit"),
						Result:  c.String("result"),
						Skip:    c.Int("skip"),
						Sort:    c.String("sort"),
						Stack:   c.String("stack"),
						Status:  c.String("status"),
					}

					return client.GetDeploys(getDeploysOptions)
				}),
			},
		},
	}
	applicationsCommand = cli.Command{
		Name:  "applications",
		Usage: "application related endpoints",
		Subcommands: []cli.Command{
			cli.Command{
				Name:  "get",
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

					getApplicationOptions := &wercker.GetApplicationOptions{Owner: owner, Name: name}
					return client.GetApplication(getApplicationOptions)
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
			os.Exit(1)
		}

		if result == nil {
			return
		}

		b, err := json.MarshalIndent(result, "", "  ")
		if err != nil {
			os.Stderr.WriteString("Unable to marshal response from the API: ")
			os.Stderr.WriteString(err.Error())
			os.Stderr.WriteString("\n")
			os.Exit(2)
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
	app.Name = "explorer"
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
		buildCommand,
		deployCommand,
		tokensCommand,
	}

	app.Run(os.Args)
}
