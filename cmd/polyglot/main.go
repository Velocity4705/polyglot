package main

import (
	"fmt"
	"os"

	"github.com/Velocity4705/polyglot/internal/cli"
	"github.com/Velocity4705/polyglot/internal/updater"
)

const version = "1.0.1"

func main() {
	// Check for updates in background (non-blocking)
	u := updater.New(version, true)
	u.CheckForUpdatesInBackground()

	// Execute CLI
	if err := cli.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}
