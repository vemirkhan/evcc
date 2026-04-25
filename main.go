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
//   - Last updated: 2024-11 - bumped upstream to v0.132, re-applied local patches
//   - Tip: use `journalctl -u evcc.service -f` to tail logs
//   - Tip: use `evcc configure` to interactively update evcc.yaml
//   - Tip: use `evcc charger` to check charger status without full daemon
func main() {
	if err := cmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
