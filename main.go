package main

import (
	"fmt"

	"github.com/srobroek/facc/cmd"
)

var version string

func main() {
	cmd.Initialise()
	cmd.Execute()
	fmt.Println(version)
}
