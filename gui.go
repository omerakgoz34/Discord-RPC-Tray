package main

import (
	"github.com/andlabs/ui"
)

func GUIsetup() {
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
	entryFormAppID := ui.NewEntry()
	entryFormAppID.SetText(Config.AppID)
	formApp.Append("App ID:", entryFormAppID, false)
	container.Append(formApp, false)

	// RPC
	groupRPC := ui.NewGroup("RPC")
	groupRPC.SetMargined(true)
	containerGroupRPC := ui.NewVerticalBox()
	containerGroupRPC.SetPadded(true)
	formGroupRPC := ui.NewForm()
	formGroupRPC.SetPadded(true)

	entryGroupRPCdetails := ui.NewEntry()
	entryGroupRPCdetails.SetText(Config.RPC.Details)
	formGroupRPC.Append("Details:", entryGroupRPCdetails, false)
	entryGroupRPCstate := ui.NewEntry()
	entryGroupRPCstate.SetText(Config.RPC.State)
	formGroupRPC.Append("State:", entryGroupRPCstate, false)
	entryGroupRPClargeImage := ui.NewEntry()
	entryGroupRPClargeImage.SetText(Config.RPC.LargeImage)
	formGroupRPC.Append("Large Image:", entryGroupRPClargeImage, false)
	entryGroupRPClargeText := ui.NewEntry()
	entryGroupRPClargeText.SetText(Config.RPC.LargeText)
	formGroupRPC.Append("Large Text:", entryGroupRPClargeText, false)
	entryGroupRPCsmallImage := ui.NewEntry()
	entryGroupRPCsmallImage.SetText(Config.RPC.SmallImage)
	formGroupRPC.Append("Small Image:", entryGroupRPCsmallImage, false)
	entryGroupRPCSmallText := ui.NewEntry()
	entryGroupRPCSmallText.SetText(Config.RPC.SmallText)
	formGroupRPC.Append("Small Text:", entryGroupRPCSmallText, false)

	containerGroupRPC.Append(formGroupRPC, false)
	groupRPC.SetChild(containerGroupRPC)
	container.Append(groupRPC, false)

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
