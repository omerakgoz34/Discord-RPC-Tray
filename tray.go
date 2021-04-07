package main

import (
	"log"

	"github.com/getlantern/systray"
	"github.com/getlantern/systray/example/icon"
)

func Tray() {
	systray.SetTemplateIcon(icon.Data, icon.Data)
	systray.SetTitle(AppName)
	systray.SetTooltip(AppName)

	// Tray Button - Show/Hide Console
	trayButtonConsole := systray.AddMenuItem(Lang["trayButtonConsoleHide"], Lang["trayButtonConsoleHideDesc"])
	trayButtonConsoleType := "hide"
	go func() {
		for {
			<-trayButtonConsole.ClickedCh
			switch trayButtonConsoleType {
			case "hide":
				ConsoleHide()
				log.Println(Lang["debugConsoleHide"])
				trayButtonConsole.SetTitle(Lang["trayButtonConsoleShow"])
				trayButtonConsole.SetTooltip(Lang["trayButtonConsoleShowDesc"])
				trayButtonConsoleType = "show"

			case "show":
				ConsoleShow()
				log.Println(Lang["debugConsoleShow"])
				trayButtonConsole.SetTitle(Lang["trayButtonConsoleHide"])
				trayButtonConsole.SetTooltip(Lang["trayButtonConsoleHideDesc"])
				trayButtonConsoleType = "hide"
			}
		}
	}()

	// Tray Button - Quit
	trayButtonQuit := systray.AddMenuItem(Lang["quit"], Lang["trayButtonQuitDesc"])
	go func() {
		<-trayButtonQuit.ClickedCh
		systray.Quit()
	}()

	log.Println(Lang["debugTrayReady"])
}
