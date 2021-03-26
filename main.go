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
	// Sync Configs
	InitConfigFile()

	// Start Core
	log.Println(">>> CORE: starting...")
	log.Println(">>> CORE: initializing the lang...")
	InitLang()
	log.Println(">>> CORE: ready.")

	// Start UI
	log.Println(">>> UI: starting...")
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
	log.Println(">>> UI: ready.")
}

// Quit ...
func Quit() {
	log.Println(">>> CORE: quitting...")
}
