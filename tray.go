package main

import (
	"log"
	"time"

	"github.com/getlantern/systray"
	"github.com/skratchdot/open-golang/open"
)

func Tray() {
	systray.SetTemplateIcon(Icon, Icon)
	systray.SetTitle(AppName)
	systray.SetTooltip(AppName)

	systray.AddMenuItem(AppName+" v"+AppVersion, "Version").Disable()
	systray.AddSeparator()

	trayButtonOpenConfig := systray.AddMenuItem("Open Config", "Opens config file")
	go func() {
		for {
			<-trayButtonOpenConfig.ClickedCh
			open.Run(ConfigDir + ConfigFileName)
			log.Println("Config opened.")
		}
	}()

	trayButtonConfigSample := systray.AddMenuItem("Open Sample Config", "Opens sample config file")
	go func() {
		for {
			<-trayButtonConfigSample.ClickedCh
			open.Run(ConfigDir + ConfigFileSampleName)
			log.Println("Sample config reloaded.")
		}
	}()

	systray.AddSeparator()

	trayButtonRPC := systray.AddMenuItem("Start RPC", "Starts RPC")
	trayButtonRPCState := true
	go func() {
		for {
			<-trayButtonRPC.ClickedCh
			switch trayButtonRPCState {
			case true:
				RPCStart()
				trayButtonRPC.SetTitle("Stop RPC")
				trayButtonRPC.SetTooltip("Stops RPC")
				systray.SetTemplateIcon(IconGreen, IconGreen)
				trayButtonRPCState = false

			case false:
				RPCStop()
				trayButtonRPC.SetTitle("Start RPC")
				trayButtonRPC.SetTooltip("Starts RPC")
				systray.SetTemplateIcon(IconRed, IconRed)
				trayButtonRPCState = true
			}
		}
	}()

	trayButtonDateNow := systray.AddMenuItem("Date Now", "Writes time.Now() to config file")
	go func() {
		for {
			<-trayButtonDateNow.ClickedCh
			Config.DateNow = time.Now()
			ConfigSave()
		}
	}()

	trayButtonReload := systray.AddMenuItem("Reload Config", "Reloads config file")
	go func() {
		for {
			<-trayButtonReload.ClickedCh
			ConfigReload()
		}
	}()

	systray.AddSeparator()

	trayButtonProjectPage := systray.AddMenuItem("Project Page", "Opens project page")
	go func() {
		<-trayButtonProjectPage.ClickedCh
		open.Run(AppProjectPage)
	}()

	trayButtonQuit := systray.AddMenuItem("Quit", "Quits the app")
	go func() {
		<-trayButtonQuit.ClickedCh
		systray.Quit()
	}()
}
