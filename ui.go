package main

import (
	"strconv"

	"github.com/getlantern/systray"
	"github.com/rivo/tview"
)

var UI *tview.Application = tview.NewApplication()

func LoopUI() {
	// UI Elements
	pages := tview.NewPages()
	lastPage := ""
	lastPageModal := ""
	formAddFirstApp := tview.NewForm()
	formAddNewApp := tview.NewForm()
	modalInvalidAppID := tview.NewModal()
	formAppSelection := tview.NewForm()
	formRPC := tview.NewForm()
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
				formRPC.SetTitle(" " + AppName + " - " + Lang["selectedApp"] + Config["selectedApp"] + " ")
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
		AddButton(Lang["save"], func() {
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
				pages.SwitchToPage(lastPage)
				lastPage = "RPC"
			}
		}).
		AddButton(Lang["cancel"], func() {
			pages.SwitchToPage(lastPage)
			lastPage = "RPC"
		})
	formAddNewApp.SetBorder(true).SetTitle(" " + Lang["formAddNewAppTitle"] + " - " + AppName + " ").SetTitleAlign(tview.AlignLeft)
	pages.AddPage("addNewApp", formAddNewApp, true, false)

	// UI Page - Invalid App ID Modal
	modalInvalidAppID = modalInvalidAppID.
		SetText(Lang["invalidAppID"]).
		AddButtons([]string{Lang["ok"]}).
		SetDoneFunc(func(buttonIndex int, buttonLabel string) {
			pages.SwitchToPage(lastPageModal)
		})
	pages.AddPage("invalidAppID", modalInvalidAppID, true, false)

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
				formRPC.SetTitle(" " + AppName + " - " + Lang["selectedApp"] + Config["selectedApp"] + " ")
			}
			pages.SwitchToPage(lastPage)
		}).
		AddButton(Lang["add"], func() {
			lastPage = "appSelection"
			pages.SwitchToPage("addNewApp")
		}).
		AddButton(Lang["remove"], func() {
			// ...
		}).
		AddButton(Lang["quit"], func() {
			systray.Quit()
		})
	formAppSelection.SetBorder(true).SetTitle(" " + AppName + " ").SetTitleAlign(tview.AlignLeft)
	pages.AddPage("appSelection", formAppSelection, true, false)

	// UI Page - RPC
	formRPC = formRPC.
		AddInputField(Lang["details"], "", 30, nil, nil).
		AddInputField(Lang["status"], "", 30, nil, nil).
		AddButton(Lang["start"], nil).
		AddButton(Lang["changeApp"], func() {
			options := []string{}
			for k := range ConfigApps {
				options = append(options, k)
			}
			formItemAppList.SetOptions(options, func(text string, index int) {
				inputAppName = text
			})
			pages.SwitchToPage("appSelection")
			lastPage = "RPC"
		}).
		AddButton(Lang["quit"], func() {
			systray.Quit()
		})
	formRPC.SetBorder(true).SetTitle(" " + Lang["selectedApp"] + Config["selectedApp"] + " ").SetTitleAlign(tview.AlignLeft)
	pages.AddPage("RPC", formRPC, true, false)

	// Check if there is no saved app
	if len(ConfigApps) < 1 {
		Config["selectedApp"] = inputAppName
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
