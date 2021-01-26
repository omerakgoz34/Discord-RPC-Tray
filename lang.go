package main

import "log"

// Lang ...
var Lang map[string]string = map[string]string{}

// InitLang sets the app language
func InitLang() {
	switch Config["lang"] {
	case "en":
		Lang = langEN
	case "tr":
		Lang = langTR
	default:
		Lang = langEN
	}
	log.Println(Lang)
}

// Lang English
var langEN map[string]string = map[string]string{
	"trayMenuQuit":     "Quit",
	"trayMenuQuitDesc": "Quit from the app",
}

// Lang Turkish
var langTR map[string]string = map[string]string{
	"trayMenuQuit":     "Çık",
	"trayMenuQuitDesc": "Uygulamadan çık",
}
