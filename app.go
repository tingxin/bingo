package main

import (
	"os"

	"github.com/tingxin/go-utility/log"
	"github.com/urfave/cli"

	"github.com/tingxin/bingo/service/data"
	"github.com/tingxin/bingo/service/resource"
)

var (
	app        *cli.App
	configPath string
)

func init() {
	// Initialise a CLI app
	app = cli.NewApp()
	app.Name = "bingo"
	app.Usage = "create a awsome bi paltform"
	app.Author = "barry.xu"
	app.Email = "friendship-119@163.com"
	app.Version = "0.0.0"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:        "c",
			Value:       "./config.yaml",
			Destination: &configPath,
			Usage:       "Path to a configuration file",
		},
	}
}

func main() {
	initCmd()
}

func initCmd() {
	// Set the CLI app commands
	app.Commands = []cli.Command{
		{
			Name:  "data",
			Usage: "start bingo data service",
			Action: func(c *cli.Context) error {
				log.INFO.Printf("start bingo data service")
				if err := runDataService(); err != nil {
					return cli.NewExitError(err.Error(), 1)
				}
				return nil
			},
		},
		{
			Name:  "resource",
			Usage: "start bingo resource service",
			Action: func(c *cli.Context) error {
				log.INFO.Printf("start bingo resource service")
				if err := runResourceService(); err != nil {
					return cli.NewExitError(err.Error(), 1)
				}
				return nil
			},
		},
	}

	// Run the CLI app
	app.Run(os.Args)
}

func runDataService() error {
	server := data.New()
	return server.Run()
}

func runResourceService() error {
	server := resource.New()
	return server.Run()
}
