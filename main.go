package main

import (
	"log"

	"github.com/getlantern/systray"
	"github.com/getlantern/systray/example/icon"
)

// AppName ...
const AppName string = "Discord RPC Tray"

func main() {
	trayExit := func() {
		log.Println("Quitting...")
	}

	systray.Run(trayReady, trayExit)
}

func trayReady() {
	systray.SetTemplateIcon(icon.Data, icon.Data)
	systray.SetTitle(AppName)
	systray.SetTooltip(AppName)

	// Çıkış
	trayQuitButton := systray.AddMenuItem("Quit", "Quit from the app")
	go func() {
		<-trayQuitButton.ClickedCh
		log.Println("Requesting Quit.")
		Quit()
	}()
}

// Quit ...
func Quit() {
	systray.Quit()
}
