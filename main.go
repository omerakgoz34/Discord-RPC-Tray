package main

import (
	"log"

	"github.com/getlantern/systray"
	"github.com/getlantern/systray/example/icon"
)

const (
	// AppName ...
	AppName = "Discord RPC Tray"
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

	// Start the Tray
	log.Println(Lang["debugTrayStarting"])
	systray.Run(ui, Quit)
}

func ui() {
	systray.SetTemplateIcon(icon.Data, icon.Data)
	systray.SetTitle(AppName)
	systray.SetTooltip(AppName)

	// Çıkış
	trayQuitButton := systray.AddMenuItem(Lang["trayMenuQuit"], Lang["trayMenuQuitDesc"])
	go func() {
		<-trayQuitButton.ClickedCh
		systray.Quit()
	}()
	log.Println(Lang["debugTrayReady"])
}

// Quit ...
func Quit() {
	log.Println(Lang["debugCoreQuitting"])
}
