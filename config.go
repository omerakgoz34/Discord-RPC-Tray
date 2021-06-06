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

	Config = struct {
		AppID   string          `json:"AppID"`
		DateNow time.Time       `json:"DateNow"`
		RPC     client.Activity `json:"RPC"`
	}{
		AppID:   "830041049794609152",
		DateNow: time.Now(),
		RPC: client.Activity{
			Details:    "Playing with RPC",
			State:      "Players: ",
			LargeImage: "discord-logo",
			LargeText:  "This is a Discord Logo",
			SmallImage: "discord-logo",
			SmallText:  "Also, this is a Discord Logo",
			Party: &client.Party{
				ID:         "343434343434343434",
				Players:    1,
				MaxPlayers: 100,
			},
			Timestamps: &client.Timestamps{},
			Buttons: []*client.Button{
				{
					Label: "Project Page",
					Url:   "https://github.com/omerakgoz34/Discord-RPC-Tray",
				},
				{
					Label: "Created by omerakgoz34",
					Url:   "https://twitter.com/omerakgoz34",
				},
			},
		},
	}
	ConfigSample = struct {
		AppID   string          `json:"AppID"`
		DateNow time.Time       `json:"DateNow"`
		RPC     client.Activity `json:"RPC"`
	}{
		AppID:   "830041049794609152",
		DateNow: time.Now(),
		RPC: client.Activity{
			Details:    "Playing with RPC",
			State:      "Players: ",
			LargeImage: "discord-logo",
			LargeText:  "This is a Discord Logo",
			SmallImage: "discord-logo",
			SmallText:  "Also, this is a Discord Logo",
			Party: &client.Party{
				ID:         "343434343434343434",
				Players:    1,
				MaxPlayers: 100,
			},
			Timestamps: &client.Timestamps{},
			Buttons: []*client.Button{
				{
					Label: "Project Page",
					Url:   "https://github.com/omerakgoz34/Discord-RPC-Tray",
				},
				{
					Label: "Created by omerakgoz34",
					Url:   "https://twitter.com/omerakgoz34",
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
	Config.DateNow = time.Now()
	log.Println("Config: ", Config)

	// Write Sample Config File
	if err := os.MkdirAll(ConfigDir, os.ModePerm); err != nil {
		log.Fatalln(err)
	}
	configFileSample, err := os.Create(ConfigDir + ConfigFileSampleName)
	if err != nil {
		log.Fatalln(err)
	}
	defer configFileSample.Close()
	configBytes, err := json.MarshalIndent(ConfigSample, "", "    ")
	if err != nil {
		log.Fatalln(err)
	}
	if _, err = configFileSample.Write(configBytes); err != nil {
		log.Fatalln(err)
	}
	if err = configFileSample.Sync(); err != nil {
		log.Fatalln(err)
	}

	ConfigSave()
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
