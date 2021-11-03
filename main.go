package main

import (
	"log"

	"github.com/andlabs/ui"
	_ "github.com/andlabs/ui/winmanifest"
	"github.com/getlantern/systray"
)

const (
	AppName        = "Discord RPC Tray"
	AppVersion     = "1.1.0"
	AppProjectPage = "https://github.com/omerakgoz34/Discord-RPC-Tray"
)

var QuitCh = make(chan bool, 1)

func main() {
	log.SetPrefix("DEBUG >>> ")

	ConfigInit()

	go func() {
		ui.Main(GUIsetup)
		QuitCh <- true
		systray.Quit()
	}()

	systray.Run(Tray, Quit)
}

func Quit() {
	RPCStop()
	<-QuitCh
	log.Println("Quitted.")
}
