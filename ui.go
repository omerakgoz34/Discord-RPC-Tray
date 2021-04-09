package main

import (
	"strconv"

	"github.com/getlantern/systray"
	"github.com/hugolgst/rich-go/client"
	"github.com/rivo/tview"
)

var (
	UI      *tview.Application = tview.NewApplication()
	FormRPC                    = tview.NewForm()
)

func rpcButtonStart() {
	RPCActive = true
	StartRPC()
	FormRPC.GetButton(0).SetLabel(Lang["stop"]).SetSelectedFunc(rpcButtonStop)
}
func rpcButtonStop() {
	RPCActive = false
	client.Logout()
	FormRPC.GetButton(0).SetLabel(Lang["start"]).SetSelectedFunc(rpcButtonStart)
}

func LoopUI() {
	// UI Elements
	pages := tview.NewPages()
	lastPage := ""
	lastPageModal := ""
	formAddFirstApp := tview.NewForm()
	formAddNewApp := tview.NewForm()
	modalInvalidAppID := tview.NewModal()
	formAppSelection := tview.NewForm()
	formRemovingApp := tview.NewForm()
	formItemAppList := tview.NewDropDown().SetFieldWidth(34).SetLabel(Lang["selectApp"])

	// Input Infos
	inputAppName := ""
	inputAppID := ""

	// UI Page - Add First App
	formAddFirstApp = formAddFirstApp.
		AddInputField(Lang["name"], "", 34, tview.InputFieldMaxLength(34), func(text string) {
			inputAppName = text
		}).
		AddInputField(Lang["id"], "", 18, tview.InputFieldMaxLength(18), func(text string) {
			inputAppID = text
		}).
		AddButton(Lang["save"], func() {
			if len(inputAppName) < 1 || len(inputAppID) < 18 {
				return
			}
			if _, err := strconv.ParseInt(inputAppID, 10, 64); err != nil {
				lastPageModal = "addFirstApp"
				pages.SwitchToPage("invalidAppID")
			} else {
				ConfigApps[inputAppName] = inputAppID
				Config["selectedApp"] = inputAppName
				ConfigSave()
				FormRPC.SetTitle(" " + AppName + " - " + Lang["selectedApp"] + Config["selectedApp"] + " ")
				pages.SwitchToPage("RPC")
				lastPage = "RPC"
			}
		}).
		AddButton(Lang["quit"], func() {
			systray.Quit()
		})
	formAddFirstApp.SetBorder(true).SetTitle(" " + Lang["formAddFirstAppTitle"] + " - " + AppName + " ").SetTitleAlign(tview.AlignLeft)
	pages.AddPage("addFirstApp", formAddFirstApp, true, false)

	// UI Page - Add New App
	formAddNewApp = formAddNewApp.
		AddInputField(Lang["name"], "", 34, tview.InputFieldMaxLength(34), func(text string) {
			inputAppName = text
		}).
		AddInputField(Lang["id"], "", 18, tview.InputFieldMaxLength(18), func(text string) {
			inputAppID = text
		}).
		AddButton(Lang["add"], func() {
			if len(inputAppName) < 1 || len(inputAppID) < 18 {
				return
			}
			if _, err := strconv.ParseInt(inputAppID, 10, 64); err != nil || len(inputAppID) < 18 {
				lastPageModal = "addNewApp"
				pages.SwitchToPage("invalidAppID")
			} else {
				ConfigApps[inputAppName] = inputAppID
				ConfigSave()
				options := []string{}
				for k := range ConfigApps {
					options = append(options, k)
				}
				formItemAppList.SetOptions(options, func(text string, index int) {
					inputAppName = text
				})
				formItemAppList.SetCurrentOption(-1)
				pages.SwitchToPage(lastPage)
				lastPage = "RPC"
			}
		}).
		AddButton(Lang["cancel"], func() {
			pages.SwitchToPage(lastPage)
			lastPage = "RPC"
		})
	formAddNewApp.SetBorder(true).SetTitle(" " + AppName + " ").SetTitleAlign(tview.AlignLeft)
	pages.AddPage("addNewApp", formAddNewApp, true, false)

	// UI Page - Invalid App ID Modal
	modalInvalidAppID = modalInvalidAppID.
		SetText(Lang["invalidAppID"]).
		AddButtons([]string{Lang["ok"]}).
		SetDoneFunc(func(buttonIndex int, buttonLabel string) {
			pages.SwitchToPage(lastPageModal)
		})
	pages.AddPage("invalidAppID", modalInvalidAppID, true, false)

	// UI Page - Removing App
	formRemovingApp = formRemovingApp.
		AddFormItem(formItemAppList).
		AddButton(Lang["remove"], func() {
			delete(ConfigApps, inputAppName)
			delete(Config, "selectedApp")
			inputAppName = ""
			ConfigSave()
			options := []string{}
			for k := range ConfigApps {
				options = append(options, k)
			}
			formItemAppList.SetOptions(options, func(text string, index int) {
				inputAppName = text
			})
			formItemAppList.SetCurrentOption(-1)
			pages.SwitchToPage(lastPage)
			lastPage = "RPC"
		}).
		AddButton(Lang["cancel"], func() {
			pages.SwitchToPage(lastPage)
			lastPage = "RPC"
		})
	formRemovingApp.SetBorder(true).SetTitle(" " + AppName + " ").SetTitleAlign(tview.AlignLeft)
	pages.AddPage("removingApp", formRemovingApp, true, false)

	// UI Page - App Selection
	formAppSelection = formAppSelection.
		AddFormItem(formItemAppList).
		AddButton(Lang["select"], func() {
			if len(inputAppName) < 1 {
				return
			}
			Config["selectedApp"] = inputAppName
			ConfigSave()
			if lastPage == "RPC" {
				FormRPC.SetTitle(" " + AppName + " - " + Lang["selectedApp"] + Config["selectedApp"] + " ")
			}
			pages.SwitchToPage(lastPage)
		}).
		AddButton(Lang["add"], func() {
			lastPage = "appSelection"
			pages.SwitchToPage("addNewApp")
			formItemAppList.SetCurrentOption(-1)
		}).
		AddButton(Lang["remove"], func() {
			lastPage = "appSelection"
			pages.SwitchToPage("removingApp")
			formItemAppList.SetCurrentOption(-1)
		}).
		AddButton(Lang["quit"], func() {
			systray.Quit()
		})
	formAppSelection.SetBorder(true).SetTitle(" " + AppName + " ").SetTitleAlign(tview.AlignLeft)
	pages.AddPage("appSelection", formAppSelection, true, false)

	// UI Page - RPC
	FormRPC = FormRPC.
		AddInputField(Lang["details"], "", 34, nil, func(text string) {
			RPCDetails = text
		}).
		AddInputField(Lang["state"], "", 34, nil, func(text string) {
			RPCState = text
		}).
		AddButton(Lang["start"], rpcButtonStart).
		AddButton(Lang["changeApp"], func() {
			RPCActive = false
			client.Logout()
			options := []string{}
			for k := range ConfigApps {
				options = append(options, k)
			}
			formItemAppList.SetOptions(options, func(text string, index int) {
				inputAppName = text
			})
			pages.SwitchToPage("appSelection")
			lastPage = "RPC"
			formItemAppList.SetCurrentOption(-1)
		}).
		AddButton(Lang["quit"], func() {
			systray.Quit()
		})
	FormRPC.SetBorder(true).SetTitle(" " + AppName + " - " + Lang["selectedApp"] + Config["selectedApp"] + " ").SetTitleAlign(tview.AlignLeft)
	pages.AddPage("RPC", FormRPC, true, false)

	// Check if there is no saved app
	if len(ConfigApps) < 1 {
		delete(Config, "selectedApp")
		ConfigSave()
		pages.SwitchToPage("addFirstApp")
		lastPage = "addFirstApp"
	} else {
		pages.SwitchToPage("RPC")
		lastPage = "RPC"
	}

	// Start UI loop
	go UI.SetRoot(pages, true).SetFocus(pages).EnableMouse(true).Run()
}
