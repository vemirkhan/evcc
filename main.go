package main

import (
	"fmt"
	"os"

	"github.com/evcc-io/evcc/cmd"
)

// main is the entry point for evcc - EV Charge Controller
// It delegates to the cmd package which uses cobra for CLI handling
func main() {
	if err := cmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
