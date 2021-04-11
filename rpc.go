package main

import (
	"log"
	"time"

	"github.com/hugolgst/rich-go/client"
)

var (
	RPCActive  = false
	RPCDetails string
	RPCState   string
)

func RPCStop() {
	RPCActive = false
	client.Logout()
	FormRPC.GetButton(0).SetLabel(Lang["start"]).SetSelectedFunc(RPCStart)
	log.Println(Lang["rpcStopped"])
}

func RPCStart() {
	RPCActive = true

	// Login to the selected app
	if err := client.Login(ConfigApps[Config["selectedApp"]]); err != nil {
		log.Println(Lang["error"] + ": " + err.Error())
		RPCStop()
		return
	}

	go func() {
		for {
			if !RPCActive {
				break
			}
			if err := client.SetActivity(client.Activity{
				Details: RPCDetails,
				State:   RPCState,
			}); err != nil {
				log.Println(Lang["error"] + ": " + err.Error())
				RPCStop()
				break
			}
			time.Sleep(time.Second * 12)
		}
	}()

	FormRPC.GetButton(0).SetLabel(Lang["stop"]).SetSelectedFunc(RPCStop)
	log.Println(Lang["rpcStarted"])
}
