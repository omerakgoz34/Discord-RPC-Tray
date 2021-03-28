package main

import (
	"log"

	"github.com/getlantern/systray"
	"github.com/getlantern/systray/example/icon" // hehe
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

	// Start the Tray
	log.Println(Lang["debugTrayStarting"])
	systray.Run(tray, Quit)
}

func tray() {
	systray.SetTemplateIcon(icon.Data, icon.Data)
	systray.SetTitle(AppName)
	systray.SetTooltip(AppName)

	// Tray Button - Hide Console
	trayButtonConsoleHide := systray.AddMenuItem(Lang["trayButtonConsoleHide"], Lang["trayButtonConsoleHideDesc"])
	go func() {
		for {
			<-trayButtonConsoleHide.ClickedCh
			ConsoleHide()
			log.Println(Lang["debugConsoleHide"])
		}
	}()

	// Tray Button - Show Console
	trayButtonConsoleShow := systray.AddMenuItem(Lang["trayButtonConsoleShow"], Lang["trayButtonConsoleShowDesc"])
	go func() {
		for {
			<-trayButtonConsoleShow.ClickedCh
			ConsoleShow()
			log.Println(Lang["debugConsoleShow"])
		}
	}()

	// Tray Button - Quit
	trayButtonQuit := systray.AddMenuItem(Lang["trayButtonQuit"], Lang["trayButtonQuitDesc"])
	go func() {
		<-trayButtonQuit.ClickedCh
		systray.Quit()
	}()

	log.Println(Lang["debugTrayReady"])
}

// Quit ...
func Quit() {
	log.Println(Lang["debugCoreQuitting"])
}
