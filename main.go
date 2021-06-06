package main

import (
	"log"

	"github.com/getlantern/systray"
)

const (
	AppName        = "Discord RPC Tray"
	AppVersion     = "1.0.0"
	AppProjectPage = "https://github.com/omerakgoz34/Discord-RPC-Tray"
)

func main() {
	log.SetPrefix("DEBUG >>> ")

	ConfigInit()

	systray.Run(Tray, Quit)
}

func Quit() {
	RPCStop()
	log.Println("Quitting...")
}
