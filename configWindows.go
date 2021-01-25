// +build windows

package main

import (
	"log"
	"os"
)

// GetConfigPath ...
func GetConfigPath() string {
	home, err := os.UserHomeDir()
	if err != nil {
		log.Fatalln("Error on getting user home folder! ", err)
	}
	return home + "\\AppData\\Roaming\\Discord RPC Tray\\"
}
