package main

import (
	"gitlab.com/tokend/hgate"
	"fmt"
	"github.com/urfave/cli"
)

const DefaultConfigPath = "./config.yaml"

var ServeCommand = cli.Command{
	Name:  "serve",
	Usage: "start proxy",
	Flags: []cli.Flag{
		cli.StringFlag{
			Name:  "config, c",
			Value: DefaultConfigPath,
		},
	},
	Action: serveAction,
}

func serveAction(c *cli.Context) error {
	app, err := hgate.NewApp(c.String("config"))
	if err != nil {
		return toCliError(fmt.Errorf("hgate initialization failed: %s", err.Error()))
	}

	app.Serve()

	return nil
}

func toCliError(err error) error {
	if err != nil {
		return cli.NewExitError(err, 1)
	}
	return nil
}
