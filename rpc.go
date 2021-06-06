package main

import (
	"log"
	"time"

	"github.com/getlantern/systray"
	"github.com/hugolgst/rich-go/client"
)

var RPCActive = false

func RPCStop() {
	RPCActive = false
	client.Logout()
	log.Println("RPC stopped.")
}

func RPCStart() {
	ConfigReload()
	RPCActive = true

	if err := client.Login(Config.AppID); err != nil {
		log.Println(err)
		systray.Quit()
	}

	log.Println("RPC started.")

	go func() {
		for {
			if !RPCActive {
				break
			}

			if err := client.SetActivity(Config.RPC); err != nil {
				log.Println(err)
				systray.Quit()
			}
			time.Sleep(time.Second * 12)
		}
	}()
}
