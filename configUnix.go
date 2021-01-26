// +build !windows

package main

import (
	"log"
	"os"
)

// GetConfigPath ...
func GetConfigPath() string {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatalln("Error on getting user home folder! ", err)
	}
	return homeDir + "/.config/" + AppName + "/"
}
