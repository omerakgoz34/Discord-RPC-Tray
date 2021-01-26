package main

import (
	"io/ioutil"
	"log"

	"github.com/getlantern/systray"
	"github.com/getlantern/systray/example/icon"
)

const (
	// AppName ...
	AppName string = "Discord RPC Tray"
)

func main() {
	// Enable Debug
	if AppDebug != true {
		log.SetOutput(ioutil.Discard)
	}

	// Start Core
	log.Println("> Starting Core <")

	// Sync Configs
	log.Println("Syncing configs...")
	InitConfigFile()
	log.Println("Initializing lang...")
	InitLang()

	// Start UI
	log.Println("> Starting UI < ")
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
}

// Quit ...
func Quit() {
	log.Println("> Quitting <")
}
