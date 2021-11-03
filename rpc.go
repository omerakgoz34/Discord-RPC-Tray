package main

import (
	"log"
	"time"

	"github.com/andlabs/ui"
	"github.com/hugolgst/rich-go/client"
)

var RPCActive = false

func RPCStop() {
	client.Logout()
	RPCActive = false
	GUIch <- "buttonStart"
	log.Println("RPC stopped.")
}

func RPCStart() {
	if err := client.Login(Config.AppID); err != nil {
		log.Println(err)
		RPCActive = false
		GUIch <- "buttonStart"
		ui.MsgBoxError(Win, "ERROR!", "Can not login to Discord RPC")
		return
	}

	go func() {
		for {
			if !RPCActive {
				break
			}

			if err := client.SetActivity(Config.RPC); err != nil {
				log.Println(err)
				RPCActive = false
				GUIch <- "buttonStart"
				ui.MsgBoxError(Win, "ERROR!", "Can not update RPC")
				return
			}
			time.Sleep(time.Second * 12)
		}
	}()

	RPCActive = true
	GUIch <- "buttonStop"
	log.Println("RPC started.")
}
