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
//   - Last updated: 2025-01 - bumped upstream to v0.133, re-applied local patches
//   - Tip: use `journalctl -u evcc.service -f` to tail logs
//   - Tip: use `evcc configure` to interactively update evcc.yaml
//   - Tip: use `evcc charger` to check charger status without full daemon
//   - Tip: use `evcc meter` to verify Shelly 3EM readings before troubleshooting
//   - Tip: use `evcc dumplogs` after a crash - logs rotate, so grab them quickly
//   - Tip: use `evcc vehicle` to check VW ID.3 SoC and range reported via We Connect
func main() {
	if err := cmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
