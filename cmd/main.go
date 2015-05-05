package main

import (
	"fmt"
	"os"

	"github.com/codegangsta/cli"
	"github.com/jarosser06/fins"
)

const finsVersion = "0.0.1"

func main() {
	app := cli.NewApp()
	app.Name = "fins"
	app.Usage = "Help troubleshoot a Chef Server Environment"
	app.Version = finsVersion
	app.Authors = []cli.Author{
		{
			Name:  "Jim Rosser",
			Email: "jarosser06@gmail.com",
		},
	}

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:   "config, c",
			Value:  "fins.json",
			Usage:  "the fins config file to use",
			EnvVar: "FINS_CONFIG",
		},
		cli.StringFlag{
			Name:   "loglevel, l",
			Value:  "error",
			Usage:  "fins log level",
			EnvVar: "FINS_LOGLEVEL",
		},
	}

	app.Commands = []cli.Command{
		{
			Name:        "outdated",
			Description: "show all cookbooks that are outdated base on supermarket",
			Usage:       "fins [global options] outdated",
			Action: func(c *cli.Context) {
				f := fins.Init(c.GlobalString("config"), c.GlobalString("loglevel"))
				f.Outdated()
			},
		},
		{
			Name:        "diff",
			Description: "shows a diff between an evironment and the chef server or between two environments",
			Usage:       "fins [global options] diff <environment> <optional_environment2>",
			Action: func(c *cli.Context) {
				f := fins.Init(c.GlobalString("config"), c.GlobalString("loglevel"))
				switch len(c.Args()) {
				case 0:
					fmt.Println("you must pass at least one environment")
					os.Exit(1)
				case 1:
					os.Exit(f.DiffServer(c.Args()[0]))
				case 2:
					os.Exit(f.DiffEnvironments(c.Args()[0], c.Args()[1]))
				}
			},
		},
	}

	cli.CommandHelpTemplate = `NAME:
   {{.Name}} - {{.Description}}
USAGE:
   {{.Usage}}{{if .Flags}}
OPTIONS:
   {{range .Flags}}{{.}}
   {{end}}{{ end }}
`

	app.Run(os.Args)
}
