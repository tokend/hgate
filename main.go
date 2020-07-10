package main

import (
	"os"

	"github.com/tokend/hgate/internal/cli"
)

func main() {
	if !cli.Run(os.Args) {
		os.Exit(1)
	}
}
