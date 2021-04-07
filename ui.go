package main

import (
	"strconv"

	"github.com/getlantern/systray"
	"github.com/rivo/tview"
)

var UI *tview.Application = tview.NewApplication()

func LoopUI() {
	pages := tview.NewPages()
	lastPage := ""

	// UI Page - Add First App
	isFirstApp := false
	inputName := ""
	inputID := ""
	formAddFirstApp := tview.NewForm().
		AddInputField(Lang["name"], "", 32, tview.InputFieldMaxLength(32), func(text string) {
			inputName = text
		}).
		AddInputField(Lang["id"], "", 32, tview.InputFieldMaxLength(18), func(text string) {
			inputID = text
		}).
		AddButton(Lang["save"], func() {
			if _, err := strconv.ParseInt(inputID, 10, 64); err != nil || len(inputID) < 18 {
				pages.SwitchToPage("invalidAppID")
			} else {
				ConfigApps[inputName] = inputID
				ConfigSave()
				pages.SwitchToPage("RPC")
				lastPage = "RPC"
			}
		}).
		AddButton(Lang["quit"], func() {
			systray.Quit()
		}).SetCancelFunc(func() {})
	formAddFirstApp.SetBorder(true).SetTitle(Lang["formAddFirstAppTitle"]).SetTitleAlign(tview.AlignLeft)
	pages.AddPage("addFirstApp", formAddFirstApp, true, false)

	// UI Page - Add New App
	formAddNewApp := tview.NewForm().
		AddInputField(Lang["name"], "", 32, tview.InputFieldMaxLength(32), func(text string) {
			inputName = text
		}).
		AddInputField(Lang["id"], "", 32, tview.InputFieldMaxLength(18), func(text string) {
			inputID = text
		}).
		AddButton(Lang["save"], func() {
			if _, err := strconv.ParseInt(inputID, 10, 64); err != nil || len(inputID) < 18 {
				pages.SwitchToPage("invalidAppID")
			} else {
				ConfigApps[inputName] = inputID
				ConfigSave()
				pages.SwitchToPage(lastPage)
			}
		}).
		AddButton(Lang["cancel"], func() {
			pages.SwitchToPage(lastPage)
		}).SetCancelFunc(func() {})
	formAddNewApp.SetBorder(true).SetTitle(Lang["formAddNewAppTitle"]).SetTitleAlign(tview.AlignLeft)
	pages.AddPage("addNewApp", formAddNewApp, true, false)

	// UI Page - Invalid App ID Modal
	modalInvalidAppID := tview.NewModal().
		SetText(Lang["invalidAppID"]).
		AddButtons([]string{Lang["ok"]}).
		SetDoneFunc(func(buttonIndex int, buttonLabel string) {
			if isFirstApp {
				pages.SwitchToPage("addFirstApp")
			} else {
				pages.SwitchToPage("addNewApp")
			}
		})
	pages.AddPage("invalidAppID", modalInvalidAppID, true, false)

	// UI Page - RPC
	formRPC := tview.NewForm().
		AddDropDown(Lang["application"], []string{"blank1", "blank2", "blank3"}, 0, nil).
		AddInputField(Lang["details"], "", 30, nil, nil).
		AddInputField(Lang["status"], "", 30, nil, nil).
		AddButton(Lang["start"], nil).
		AddButton(Lang["quit"], func() {
			systray.Quit()
		})
	formRPC.SetBorder(true).SetTitle(AppName).SetTitleAlign(tview.AlignLeft)
	pages.AddPage("RPC", formRPC, true, false)

	if len(ConfigApps) < 1 {
		isFirstApp = true
		pages.SwitchToPage("addFirstApp")
	} else {
		isFirstApp = false
		pages.SwitchToPage("RPC")
	}

	go UI.SetRoot(pages, true).SetFocus(pages).EnableMouse(true).Run()
}
