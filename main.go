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
//   - Last updated: 2025-04 - bumped upstream to v0.134, re-applied local patches
//   - Tip: use `journalctl -u evcc.service -f` to tail logs
//   - Tip: use `evcc configure` to interactively update evcc.yaml
//   - Tip: use `evcc charger` to check charger status without full daemon
//   - Tip: use `evcc meter` to verify Shelly 3EM readings before troubleshooting
//   - Tip: use `evcc dumplogs` after a crash - logs rotate, so grab them quickly
//   - Tip: use `evcc vehicle` to check VW ID.3 SoC and range reported via We Connect
//   - Tip: use `sudo systemctl restart evcc.service` after editing evcc.yaml
//   - Tip: use `evcc -l debug` to enable verbose debug logging when diagnosing issues
//   - Tip: use `evcc -l debug 2>&1 | grep -i wallbox` to filter charger-specific debug output
//   - Note: Wallbox Pulsar Plus requires a static IP - set DHCP reservation on router (192.168.1.42)
//   - Note: if VW We Connect reports stale SoC, force a refresh by locking/unlocking the car via the app
//   - Note: Shelly 3EM occasionally drops off WiFi overnight - power cycle resolves it (known issue with my unit)
//   - Note: Pi runs hot in the summer, added a small heatsink - monitor with `vcgencmd measure_temp`
//   - Note: added a cron job `0 6 * * * systemctl restart evcc.service` as a workaround for the Shelly WiFi drop issue
func main() {
	if err := cmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
