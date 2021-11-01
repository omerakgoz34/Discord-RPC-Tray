package main

import (
	"github.com/andlabs/ui"
)

func GUIsetup() {
	// Window
	win := ui.NewWindow(AppName, 600, 300, false)
	win.OnClosing(func(*ui.Window) bool {
		win.Destroy()
		ui.Quit()
		return false
	})
	win.SetMargined(true)
	ui.OnShouldQuit(func() bool {
		win.Destroy()
		return true
	})

	// Main Container
	container := ui.NewVerticalBox()
	container.SetPadded(true)

	// App
	formApp := ui.NewForm()
	formApp.SetPadded(true)
	container.Append(formApp, false)
	entryFormAppID := ui.NewEntry()
	entryFormAppID.SetText("AppID")
	formApp.Append("App ID:", entryFormAppID, false)

	// RPC
	entryRPCdetails := ui.NewEntry()
	entryRPCdetails.SetText("Details")
	formApp.Append("Details:", entryRPCdetails, false)
	entryRPCstate := ui.NewEntry()
	entryRPCstate.SetText("State")
	formApp.Append("State:", entryRPCstate, false)

	// RPC - Images
	entryRPCimagesLarge := ui.NewEntry()
	entryRPCimagesLarge.SetText("Large Image")
	formApp.Append("Large Image:", entryRPCimagesLarge, false)
	entryRPCimagesLargeText := ui.NewEntry()
	entryRPCimagesLargeText.SetText("Large Image Text")
	formApp.Append("Large Image Text:", entryRPCimagesLargeText, false)

	entryRPCimagesSmall := ui.NewEntry()
	entryRPCimagesSmall.SetText("Small Image")
	formApp.Append("Small Image:", entryRPCimagesSmall, false)
	entryRPCimagesSmallText := ui.NewEntry()
	entryRPCimagesSmallText.SetText("Small Image Text")
	formApp.Append("Small Image Text:", entryRPCimagesSmallText, false)

	// RPC - Party
	entryRPCpartyID := ui.NewEntry()
	entryRPCpartyID.SetText("PartyID")
	formApp.Append("Party ID:", entryRPCpartyID, false)
	entryRPCpartyPlayers := ui.NewEntry()
	entryRPCpartyPlayers.SetText("Players")
	formApp.Append("Party Players:", entryRPCpartyPlayers, false)
	entryRPCpartyMaxPLayers := ui.NewEntry()
	entryRPCpartyMaxPLayers.SetText("MaxPlayers")
	formApp.Append("Party Max Players:", entryRPCpartyMaxPLayers, false)

	// RPC - Timestamps
	entryRPCtimestampsStart := ui.NewEntry()
	entryRPCtimestampsStart.SetText("Start")
	formApp.Append("Timestamps Start:", entryRPCtimestampsStart, false)
	entryRPCtimestampsEnd := ui.NewEntry()
	entryRPCtimestampsEnd.SetText("End")
	formApp.Append("Timestamps End:", entryRPCtimestampsEnd, false)

	// RPC - Buttons
	entryRPCbuttonsFirstLabel := ui.NewEntry()
	entryRPCbuttonsFirstLabel.SetText("First Button Label")
	formApp.Append("First Button Label:", entryRPCbuttonsFirstLabel, false)
	entryRPCbuttonsFirstURL := ui.NewEntry()
	entryRPCbuttonsFirstURL.SetText("First Button URL")
	formApp.Append("First Button URL:", entryRPCbuttonsFirstURL, false)

	entryRPCbuttonsSecondLabel := ui.NewEntry()
	entryRPCbuttonsSecondLabel.SetText("Second Button Label")
	formApp.Append("Second Button Label:", entryRPCbuttonsSecondLabel, false)
	entryRPCbuttonsSecondURL := ui.NewEntry()
	entryRPCbuttonsSecondURL.SetText("Second Button URL")
	formApp.Append("Second Button URL:", entryRPCbuttonsSecondURL, false)

	// Final UI
	win.SetChild(container)
	win.Show()
	// go GUIloop()
}

func GUIloop() {
	for {
		// Update the UI using the QueueMain function
		ui.QueueMain(func() {
		})
	}
}
