package main

import (
	"io/ioutil"
	"log"
	"os"

	"github.com/getlantern/systray"
	"github.com/getlantern/systray/example/icon"
)

const (
	// AppName ...
	AppName string = "Discord RPC Tray"
)

func main() {
	if os.Getenv("DiscordRPCTrayDebug") == "true" {
		log.SetOutput(ioutil.Discard)
	}
	log.Println("Starting...")
	log.Println("Syncing config...")
	InitConfigFile()
	InitLang()
	defer ConfigFile.Close()

	log.Println("Starting tray...")
	systray.Run(tray, Quit)
}

func tray() {
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
	log.Println("Quitting...")
}
