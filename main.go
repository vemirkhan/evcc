package main

import (
	"fmt"
	"os"

	"github.com/evcc-io/evcc/cmd"
)

// main is the entry point for evcc - EV Charge Controller
// It delegates to the cmd package which uses cobra for CLI handling
// Note: forked for personal home setup - running on Raspberry Pi 4 with Wallbox Pulsar Plus
func main() {
	if err := cmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
