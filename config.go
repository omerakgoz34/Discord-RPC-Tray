package main

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"os"
	"runtime"
	"time"

	"github.com/hugolgst/rich-go/client"
)

var (
	ConfigFileName       = "config.json"
	ConfigFileSampleName = "configSample.json"
	ConfigDir            string
	now                  = time.Now()

	Config = struct {
		AppID string          `json:"AppID"`
		RPC   client.Activity `json:"RPC"`
	}{
		AppID: "830041049794609152",
	}
	ConfigSample = struct {
		AppID string          `json:"AppID"`
		RPC   client.Activity `json:"RPC"`
	}{
		AppID: "830041049794609152",
		RPC: client.Activity{
			Details:    "Playing with RPC",
			State:      "Players: ",
			LargeImage: "discord-logo",
			LargeText:  "This is a Discord Logo",
			SmallImage: "discord-logo",
			SmallText:  "Also, this is a Discord Logo",
			Party: &client.Party{
				ID:         "123456789123456789",
				Players:    1,
				MaxPlayers: 100,
			},
			Timestamps: &client.Timestamps{
				Start: &now,
				End:   &now,
			},
			Buttons: []*client.Button{
				{
					Label: "Project Page",
					Url:   "https://github.com/omerakgoz34/Discord-RPC-Tray",
				},
				{
					Label: "Created by omerakgoz34",
					Url:   "https://linktr.ee/omerakgoz34",
				},
			},
		},
	}
)

func ConfigInit() {
	// Get HOME Dir
	if runtime.GOOS == "windows" {
		homeDir, err := os.UserHomeDir()
		if err != nil {
			log.Fatalln(err)
		}
		ConfigDir = homeDir + "\\AppData\\Roaming\\" + AppName + "\\"
	} else {
		homeDir, err := os.UserHomeDir()
		if err != nil {
			log.Fatalln(err)
		}
		ConfigDir = homeDir + "/.config/" + AppName + "/"
	}

	// Load Config
	configFileBuf := bytes.NewBuffer(nil)
	var configFile *os.File
	configFile, err := os.Open(ConfigDir + ConfigFileName)
	if err != nil {
		if !os.IsExist(err) {
			// Write Sample Config File
			if err := os.MkdirAll(ConfigDir, os.ModePerm); err != nil {
				log.Fatalln(err)
			}
			configFile, err := os.Create(ConfigDir + ConfigFileName)
			if err != nil {
				log.Fatalln(err)
			}
			configBytes, err := json.MarshalIndent(ConfigSample, "", "    ")
			if err != nil {
				log.Fatalln(err)
			}
			if _, err = configFile.Write(configBytes); err != nil {
				log.Fatalln(err)
			}
			if err = configFile.Sync(); err != nil {
				log.Fatalln(err)
			}
			configFile.Close()
			ConfigReload()
		} else {
			log.Fatalln(err)
		}
	} else {
		io.Copy(configFileBuf, configFile)
		json.Unmarshal(configFileBuf.Bytes(), &Config)
		configFile.Close()
	}
	log.Println("Config: ", Config)
}

func ConfigReload() {
	configFileBuf := bytes.NewBuffer(nil)
	configFile, err := os.Open(ConfigDir + ConfigFileName)
	if err != nil {
		if !os.IsExist(err) {
		} else {
			log.Fatalln(err)
		}
	} else {
		io.Copy(configFileBuf, configFile)
		json.Unmarshal(configFileBuf.Bytes(), &Config)
	}

	configFile.Close()
	log.Println("Config reloaded.")
}

func ConfigSave() {
	if err := os.MkdirAll(ConfigDir, os.ModePerm); err != nil {
		log.Fatalln(err)
	}
	configFile, err := os.Create(ConfigDir + ConfigFileName)
	if err != nil {
		log.Fatalln(err)
	}
	defer configFile.Close()
	configBytes, err := json.MarshalIndent(Config, "", "    ")
	if err != nil {
		log.Fatalln(err)
	}
	if _, err = configFile.Write(configBytes); err != nil {
		log.Fatalln(err)
	}
	if err = configFile.Sync(); err != nil {
		log.Fatalln(err)
	}

	log.Println("Config saved.")
}
