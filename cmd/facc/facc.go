package main

import (
	"os"

	"github.com/srobroek/facc/internal/facc/cli"
)

//var settings = require.initSettings()

func main() {

	// Execute cobra
	if err := cli.Execute(); err != nil {
		//pterm.Error.PrintOnError(err)
		os.Exit(1)

	}

}
