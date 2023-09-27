package main

import (
	"github.com/nshntarora/pop/soda/cmd"
)

func main() {
	cmd.RootCmd.Use = "soda"
	cmd.Execute()
}
