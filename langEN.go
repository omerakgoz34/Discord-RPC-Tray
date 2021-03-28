package main

// Lang English
var LangEN = map[string]string{
	// Errors
	"errorGettingHomeFolder":    "ERROR: couldn't get user's home folder path!",
	"errorOpeningConfigFile":    "ERROR: couldn't open the config file!",
	"errorCreatingConfigFolder": "ERROR: couldn't create config folder!",
	"errorMarshalingConfigData": "ERROR: couldn't marshal the config data!",
	"errorWritingConfigFile":    "ERROR: couldn't write to the config file!",
	"errorSyncingConfigFile":    "ERROR: couldn't sync the config file!",

	// Debug - CONFIG
	"debugConfigSaving": "CONFIG: saving...",
	"debugConfigSaved":  "CONFIG: saved.",

	// Debug - CORE
	"debugCoreInitLang": "CORE: initialized the lang.",
	"debugCoreReady":    "CORE: ready.",
	"debugCoreQuitting": "CORE: quitting...",
	"debugConsoleHide":  "CORE: hiding console.",
	"debugConsoleShow":  "CORE: showing console.",

	// Debug - TRAY
	"debugTrayStarting": "TRAY: starting...",
	"debugTrayReady":    "TRAY: ready.",

	// Tray Menu
	"trayButtonConsoleHide":     "Hide Console",
	"trayButtonConsoleHideDesc": "Hide console of the app",
	"trayButtonConsoleShow":     "Show Console",
	"trayButtonConsoleShowDesc": "Show console of the app",
	"trayButtonQuit":            "Quit",
	"trayButtonQuitDesc":        "Quit the app",
}
