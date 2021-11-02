package main

import (
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/andlabs/ui"
	"github.com/atotto/clipboard"
	"github.com/getlantern/systray"
	"github.com/hugolgst/rich-go/client"
	"github.com/skratchdot/open-golang/open"
)

var GUIch = make(chan string)

var timeLayout = "2006-01-02 15:04:05.999999999 -0700 MST"

var Win *ui.Window
var buttonStartStop *ui.Button

func GUIsetup() {
	// Window
	Win = ui.NewWindow(AppName, 600, 300, false)
	Win.OnClosing(func(*ui.Window) bool {
		GUIch <- "quit"
		return false
	})
	Win.SetMargined(true)
	ui.OnShouldQuit(func() bool {
		GUIch <- "quit"
		return true
	})

	// Main Container
	container := ui.NewVerticalBox()
	container.SetPadded(true)
	Win.SetChild(container)

	// App
	formApp := ui.NewForm()
	formApp.SetPadded(true)
	container.Append(formApp, false)
	entryFormAppID := ui.NewEntry()
	entryFormAppID.SetText(Config.AppID)
	formApp.Append("App ID:", entryFormAppID, false)

	// RPC
	entryRPCdetails := ui.NewEntry()
	entryRPCdetails.SetText(Config.RPC.Details)
	formApp.Append("Details:", entryRPCdetails, false)
	entryRPCstate := ui.NewEntry()
	entryRPCstate.SetText(Config.RPC.State)
	formApp.Append("State:", entryRPCstate, false)

	// RPC - Images
	entryRPCimagesLarge := ui.NewEntry()
	entryRPCimagesLarge.SetText(Config.RPC.LargeImage)
	formApp.Append("Large Image:", entryRPCimagesLarge, false)
	entryRPCimagesLargeText := ui.NewEntry()
	entryRPCimagesLargeText.SetText(Config.RPC.LargeText)
	formApp.Append("Large Image Text:", entryRPCimagesLargeText, false)

	entryRPCimagesSmall := ui.NewEntry()
	entryRPCimagesSmall.SetText(Config.RPC.SmallImage)
	formApp.Append("Small Image:", entryRPCimagesSmall, false)
	entryRPCimagesSmallText := ui.NewEntry()
	entryRPCimagesSmallText.SetText(Config.RPC.SmallText)
	formApp.Append("Small Image Text:", entryRPCimagesSmallText, false)

	// RPC - Party
	entryRPCpartyID := ui.NewEntry()
	entryRPCpartyID.SetText(Config.RPC.Party.ID)
	formApp.Append("Party ID:", entryRPCpartyID, false)
	entryRPCpartyPlayers := ui.NewEntry()
	entryRPCpartyPlayers.SetText(strconv.Itoa(Config.RPC.Party.Players))
	formApp.Append("Party Players:", entryRPCpartyPlayers, false)
	entryRPCpartyMaxPLayers := ui.NewEntry()
	entryRPCpartyMaxPLayers.SetText(strconv.Itoa(Config.RPC.Party.MaxPlayers))
	formApp.Append("Party Max Players:", entryRPCpartyMaxPLayers, false)

	// RPC - Timestamps
	entryRPCtimestampsStart := ui.NewEntry()
	// entryRPCtimestampsStart.SetText(Config.RPC.Timestamps.Start.Local().Format(timeLayout))
	formApp.Append("Start Timestamp:", entryRPCtimestampsStart, false)
	entryRPCtimestampsEnd := ui.NewEntry()
	// entryRPCtimestampsEnd.SetText(Config.RPC.Timestamps.End.Local().Format(timeLayout))
	formApp.Append("End Timestamp:", entryRPCtimestampsEnd, false)

	if Config.RPC.Timestamps.Start == nil {
		entryRPCtimestampsStart.SetText("")
	} else {
		entryRPCtimestampsStart.SetText(Config.RPC.Timestamps.Start.Local().Format(timeLayout))
	}
	if Config.RPC.Timestamps.End == nil {
		entryRPCtimestampsEnd.SetText("")
		log.Println("asdasd")
	} else {
		entryRPCtimestampsEnd.SetText(Config.RPC.Timestamps.End.Local().Format(timeLayout))
	}

	// RPC - Buttons
	entryRPCbuttonsFirstLabel := ui.NewEntry()
	entryRPCbuttonsFirstLabel.SetText(Config.RPC.Buttons[0].Label)
	formApp.Append("First Button Label:", entryRPCbuttonsFirstLabel, false)
	entryRPCbuttonsFirstURL := ui.NewEntry()
	entryRPCbuttonsFirstURL.SetText(Config.RPC.Buttons[0].Url)
	formApp.Append("First Button URL:", entryRPCbuttonsFirstURL, false)

	entryRPCbuttonsSecondLabel := ui.NewEntry()
	entryRPCbuttonsSecondLabel.SetText(Config.RPC.Buttons[1].Label)
	formApp.Append("Second Button Label:", entryRPCbuttonsSecondLabel, false)
	entryRPCbuttonsSecondURL := ui.NewEntry()
	entryRPCbuttonsSecondURL.SetText(Config.RPC.Buttons[1].Url)
	formApp.Append("Second Button URL:", entryRPCbuttonsSecondURL, false)

	// Control Buttons
	containerButtons := ui.NewHorizontalBox()
	containerButtons.SetPadded(true)

	buttonStartStop = ui.NewButton("   Start RPC   ")
	buttonStartStop.OnClicked(func(*ui.Button) {
		if !RPCActive {
			buttonStartStop.Disable()
			RPCStart()
		} else if RPCActive {
			buttonStartStop.Disable()
			RPCStop()
		}
	})

	buttonTimeNow := ui.NewButton("   Copy Current Time   ")
	buttonTimeNow.OnClicked(func(*ui.Button) {
		buttonTimeNow.Disable()
		clipboard.WriteAll(time.Now().Local().Format(timeLayout))
		buttonTimeNow.Enable()
		log.Println("Current time copied to the clipboard.")
	})

	buttonSaveConfig := ui.NewButton("   Save Config   ")
	buttonSaveConfig.OnClicked(func(*ui.Button) {
		buttonSaveConfig.Disable()

		// IDs
		_, err := strconv.Atoi(strings.TrimSpace(entryFormAppID.Text()))
		if err != nil {
			buttonSaveConfig.Enable()
			ui.MsgBoxError(Win, "ERROR!", "App ID field has to be number.")
			return
		} else if len(strings.TrimSpace(entryFormAppID.Text())) < 18 {
			buttonSaveConfig.Enable()
			ui.MsgBoxError(Win, "ERROR!", "App ID is not valid.")
			return
		} else {
			Config.AppID = strings.TrimSpace(entryFormAppID.Text())
		}

		if len(strings.TrimSpace(entryRPCpartyID.Text())) > 0 {
			_, err = strconv.Atoi(strings.TrimSpace(entryRPCpartyID.Text()))
			if err != nil {
				buttonSaveConfig.Enable()
				ui.MsgBoxError(Win, "ERROR!", "Party ID field has to be number.")
				return
			} else if len(strings.TrimSpace(entryRPCpartyID.Text())) < 18 {
				buttonSaveConfig.Enable()
				ui.MsgBoxError(Win, "ERROR!", "Party ID is not valid.")
				return
			} else {
				Config.RPC.Party.ID = entryRPCpartyID.Text()
			}
		} else {
			Config.RPC.Party.ID = ""
		}

		// Party Players
		if len(strings.TrimSpace(entryRPCpartyPlayers.Text())) > 0 {
			numP, err := strconv.Atoi(strings.TrimSpace(entryRPCpartyPlayers.Text()))
			if err != nil {
				buttonSaveConfig.Enable()
				ui.MsgBoxError(Win, "ERROR!", "Party Players field has to be number.")
				return
			} else {
				Config.RPC.Party.Players = numP
			}
		} else {
			Config.RPC.Party.Players = 0
		}

		if len(strings.TrimSpace(entryRPCpartyMaxPLayers.Text())) > 0 {
			numX, err := strconv.Atoi(strings.TrimSpace(entryRPCpartyMaxPLayers.Text()))
			if err != nil {
				buttonSaveConfig.Enable()
				ui.MsgBoxError(Win, "ERROR!", "Party Max Players field has to be number.")
				return
			} else {
				Config.RPC.Party.MaxPlayers = numX
			}
		} else {
			Config.RPC.Party.MaxPlayers = 0
		}

		// Timestamps
		if len(strings.TrimSpace(entryRPCtimestampsStart.Text())) > 0 {
			tmS, err := time.Parse(timeLayout, strings.TrimSpace(entryRPCtimestampsStart.Text()))
			if err != nil {
				buttonSaveConfig.Enable()
				ui.MsgBoxError(Win, "ERROR!", "Start Timestamp is not valid.")
				return
			} else {
				Config.RPC.Timestamps.Start = &tmS
			}
		} else {
			Config.RPC.Timestamps.Start = nil
		}
		if len(strings.TrimSpace(entryRPCtimestampsEnd.Text())) > 0 {
			tmE, err := time.Parse(timeLayout, strings.TrimSpace(entryRPCtimestampsEnd.Text()))
			if err != nil {
				buttonSaveConfig.Enable()
				ui.MsgBoxError(Win, "ERROR!", "End Timestamp is not valid.")
				return
			} else {
				Config.RPC.Timestamps.End = &tmE
			}
		} else {
			Config.RPC.Timestamps.End = nil
		}

		// Texts
		Config.RPC.Details = entryRPCdetails.Text()
		Config.RPC.State = entryRPCstate.Text()
		Config.RPC.LargeImage = entryRPCimagesLarge.Text()
		Config.RPC.LargeText = entryRPCimagesLargeText.Text()
		Config.RPC.SmallImage = entryRPCimagesSmall.Text()
		Config.RPC.SmallText = entryRPCimagesSmallText.Text()
		Config.RPC.Buttons[0].Label = entryRPCbuttonsFirstLabel.Text()
		Config.RPC.Buttons[0].Url = entryRPCbuttonsFirstURL.Text()
		Config.RPC.Buttons[1].Label = entryRPCbuttonsSecondLabel.Text()
		Config.RPC.Buttons[1].Url = entryRPCbuttonsSecondURL.Text()
		ConfigSave()

		if RPCActive {
			if err := client.SetActivity(Config.RPC); err != nil {
				log.Println(err)
				RPCActive = false
				GUIch <- "buttonStart"
				ui.MsgBoxError(Win, "ERROR!", "Can't update RPC")
			}
		}

		buttonSaveConfig.Enable()
	})

	buttonReloadConfig := ui.NewButton("   Reload Config   ")
	buttonReloadConfig.OnClicked(func(*ui.Button) {
		buttonReloadConfig.Disable()

		ConfigReload()
		entryFormAppID.SetText(Config.AppID)
		entryRPCdetails.SetText(Config.RPC.Details)
		entryRPCstate.SetText(Config.RPC.State)
		entryRPCimagesLarge.SetText(Config.RPC.LargeImage)
		entryRPCimagesLargeText.SetText(Config.RPC.LargeText)
		entryRPCimagesSmall.SetText(Config.RPC.SmallImage)
		entryRPCimagesSmallText.SetText(Config.RPC.SmallText)
		entryRPCpartyID.SetText(Config.RPC.Party.ID)
		entryRPCpartyPlayers.SetText(strconv.Itoa(Config.RPC.Party.Players))
		entryRPCpartyMaxPLayers.SetText(strconv.Itoa(Config.RPC.Party.MaxPlayers))
		if Config.RPC.Timestamps.Start == nil {
			entryRPCtimestampsStart.SetText("")
		} else {
			entryRPCtimestampsStart.SetText(Config.RPC.Timestamps.Start.Local().Format(timeLayout))
		}
		if Config.RPC.Timestamps.End == nil {
			entryRPCtimestampsEnd.SetText("")
		} else {
			entryRPCtimestampsEnd.SetText(Config.RPC.Timestamps.End.Local().Format(timeLayout))
		}
		entryRPCbuttonsFirstLabel.SetText(Config.RPC.Buttons[0].Label)
		entryRPCbuttonsFirstURL.SetText(Config.RPC.Buttons[0].Url)
		entryRPCbuttonsSecondLabel.SetText(Config.RPC.Buttons[1].Label)
		entryRPCbuttonsSecondURL.SetText(Config.RPC.Buttons[1].Url)

		if RPCActive {
			if err := client.SetActivity(Config.RPC); err != nil {
				log.Println(err)
				RPCActive = false
				GUIch <- "buttonStart"
				ui.MsgBoxError(Win, "ERROR!", "Can't update RPC")
			}
		}

		buttonReloadConfig.Enable()
	})

	buttonOpenConfigFile := ui.NewButton("   Open Config File   ")
	buttonOpenConfigFile.OnClicked(func(*ui.Button) {
		buttonOpenConfigFile.Disable()
		if err := open.Run(ConfigDir + ConfigFileName); err != nil {
			log.Println("Can't open config file.")
			ui.MsgBoxError(Win, "ERROR!", "Can't open config file.")
		} else {
			buttonOpenConfigFile.Enable()
			log.Println("Config file opened.")
		}
	})

	buttonHideWin := ui.NewButton("   Hide Window   ")
	buttonHideWin.OnClicked(func(*ui.Button) {
		GUIch <- "windowHide"
	})

	buttonQuit := ui.NewButton("   Quit   ")
	buttonQuit.OnClicked(func(*ui.Button) {
		GUIch <- "quit"
	})

	containerButtons.Append(buttonStartStop, false)
	containerButtons.Append(buttonTimeNow, false)
	containerButtons.Append(buttonSaveConfig, false)
	containerButtons.Append(buttonReloadConfig, false)
	containerButtons.Append(buttonOpenConfigFile, false)
	containerButtons.Append(buttonHideWin, false)
	containerButtons.Append(buttonQuit, false)
	container.Append(containerButtons, false)

	Win.Show()

	if ConfigCurropted {
		ui.MsgBox(Win, "Warning!", "Config file was curropted and created new one.")
		ConfigCurropted = false
	}

	go GUIloop()
}

func GUIloop() {
	for {
		act := <-GUIch
		switch act {
		case "quit":
			ui.QueueMain(func() {
				Win.Destroy()
				ui.Quit()
			})

		case "buttonStart":
			ui.QueueMain(func() {
				systray.SetTemplateIcon(IconRed, IconRed)
				buttonStartStop.SetText("   Start RPC   ")
				buttonStartStop.Enable()
			})

		case "buttonStop":
			ui.QueueMain(func() {
				systray.SetTemplateIcon(IconGreen, IconGreen)
				buttonStartStop.SetText("   Stop RPC   ")
				buttonStartStop.Enable()
			})

		case "windowHide":
			ui.QueueMain(func() {
				TrayButtonShowHide.Disable()
				Win.Hide()
				Win.Disable()
				TrayButtonShowHide.SetTitle("Show Window")
				TrayButtonShowHide.SetTooltip("Shows Window")
				TrayButtonShowHide.Enable()
				StatusTrayButtonShowHide = true
				log.Println("Window hide.")
			})

		case "windowShow":
			ui.QueueMain(func() {
				TrayButtonShowHide.Disable()
				Win.Enable()
				Win.Show()
				TrayButtonShowHide.SetTitle("Hide Window")
				TrayButtonShowHide.SetTooltip("Hides Window")
				TrayButtonShowHide.Enable()
				StatusTrayButtonShowHide = false
				log.Println("Window show.")
			})
		}
	}
}
