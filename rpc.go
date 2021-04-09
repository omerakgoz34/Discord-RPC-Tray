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

func StartRPC() {
	if err := client.Login(ConfigApps[Config["selectedApp"]]); err != nil {
		log.Fatalln(err)
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
				log.Println(err)
				RPCActive = false
				client.Logout()
			}
			time.Sleep(time.Second * 13)
		}
	}()
}
