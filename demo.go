package main

import "os"

func isDemoMode() bool {
	return os.Getenv("DEMO_ENV") == "1" || os.Getenv("TUI_HUB_DEMO") == "1"
}
