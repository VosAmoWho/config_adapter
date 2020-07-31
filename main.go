package main

import (
	"log"
	"os"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()

	app.Commands = []*cli.Command{
		{
			Name:     "config",
			Usage:    "modify config files",
			Category: "Config",
			Subcommands: []*cli.Command{
				{
					Name:   "svp",
					Usage:  "modify output save path",
					Action: ModifySavePath,
				},
				{
					Name:   "pkg",
					Usage:  "modify package of testing",
					Action: ModifyPackage,
				},
			},
		},
	}

	app.Name = "config modifier"
	app.Usage = "for changing config file of mobilperf"
	app.Description = "debug version"
	app.Version = "0.1.0"

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
