package main

import (
	"github.com/urfave/cli"
	"os"
)

func main() {
	cmd := cli.NewApp()

	cmd.Commands = []cli.Command{
		ServeCommand,
	}

	cmd.Run(os.Args)
}
