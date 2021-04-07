package main

import (
	"log"

	"github.com/getlantern/systray" // hehe
)

const (
	// AppName ...
	AppName    = "Discord RPC Tray"
	AppVersion = "v0.0.0"
)

func main() {
	// Logging Msg Prefix
	log.SetPrefix("DEBUG >>> ")

	// Start the Core
	log.Println("CORE: starting...")
	InitConfigFile()
	InitLang()
	log.Println(Lang["debugCoreInitLang"])
	log.Println(Lang["debugCoreReady"])

	// Start the UI
	log.Println(Lang["debugUIStarting"])
	LoopUI()
	// Start the Tray
	log.Println(Lang["debugTrayStarting"])
	systray.Run(Tray, Quit)
}

// Quit ...
func Quit() {
	UI.Stop()
	log.Println(Lang["debugCoreQuitting"])
}
