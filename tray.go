package main

import (
	"github.com/getlantern/systray"
	"github.com/skratchdot/open-golang/open"
)

var StatusTrayButtonShowHide = false
var TrayButtonShowHide *systray.MenuItem

func Tray() {
	systray.SetTemplateIcon(Icon, Icon)
	systray.SetTitle(AppName)
	systray.SetTooltip(AppName)

	systray.AddMenuItem(AppName+" v"+AppVersion, "Version").Disable()
	systray.AddSeparator()

	TrayButtonShowHide = systray.AddMenuItem("Hide Window", "Hides Window")
	go func() {
		for {
			<-TrayButtonShowHide.ClickedCh
			if !StatusTrayButtonShowHide {
				GUIch <- "windowHide"
			} else if StatusTrayButtonShowHide {
				GUIch <- "windowShow"
			}
		}
	}()

	trayButtonProjectPage := systray.AddMenuItem("Open Project Page", "Opens project page")
	go func() {
		for {
			<-trayButtonProjectPage.ClickedCh
			open.Run(AppProjectPage)
		}
	}()

	trayButtonQuit := systray.AddMenuItem("Quit", "Quits the app")
	go func() {
		<-trayButtonQuit.ClickedCh
		GUIch <- "quit"
	}()
}
