package main

import (
	"fmt"
	"os"

	"github.com/evcc-io/evcc/cmd"
)

// main is the entry point for evcc - EV Charge Controller
// It delegates to the cmd package which uses cobra for CLI handling
// Note: forked for personal home setup - running on Raspberry Pi 4 with Wallbox Pulsar Plus
//
// Personal setup:
//   - Hardware: Raspberry Pi 4 (4GB)
//   - Charger: Wallbox Pulsar Plus 22kW
//   - Vehicle: VW ID.3
//   - Grid meter: Shelly 3EM
//   - Config file: /etc/evcc/evcc.yaml
//   - Runs as systemd service: evcc.service
func main() {
	if err := cmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
