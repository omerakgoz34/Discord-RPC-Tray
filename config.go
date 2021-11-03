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

type config struct {
	AppID string          `json:"AppID"`
	RPC   client.Activity `json:"RPC"`
}

var (
	ConfigFileName  = "config.json"
	ConfigDir       string
	nowSample       = time.Now()
	ConfigCurropted = false

	Config       config
	ConfigSample = config{
		AppID: "830041049794609152",
		RPC: client.Activity{
			Details:    "Playing with RPC",
			State:      "Players: ",
			LargeImage: "discord-logo",
			LargeText:  "This is a Discord Logo",
			SmallImage: "discord-logo",
			SmallText:  "Also, this is a Discord Logo",
			Party: &client.Party{
				ID:         "111111111111111111",
				Players:    1,
				MaxPlayers: 10,
			},
			Timestamps: &client.Timestamps{
				Start: &nowSample,
				End:   nil,
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

	ConfigReload()
}

func ConfigReload() {
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
			log.Println("New config file created.")
			ConfigReload()
		} else {
			log.Fatalln(err)
		}
	} else {
		io.Copy(configFileBuf, configFile)
		if err := json.Unmarshal(configFileBuf.Bytes(), &Config); err != nil {
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
			log.Println("New config file created.")
			ConfigCurropted = true
			ConfigReload()
		}
		configFile.Close()
	}

	if len(Config.AppID) < 18 {
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
		log.Println("New config file created.")
		ConfigCurropted = true
		ConfigReload()
	}
	log.Println("Config: ", Config)
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
