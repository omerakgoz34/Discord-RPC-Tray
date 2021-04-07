package main

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"os"
	"runtime"
)

var (
	// ConfigFileName ...
	ConfigFileName = "config.json"
	// ConfigAppsFileName ...
	ConfigAppsFileName = "configApps.json"
	// ConfigDir ...
	ConfigDir string

	// ConfigDefault ...
	ConfigDefault = map[string]string{
		"debug": "off",
		"lang":  "en",
	}
	// Config - Settings of the app
	Config = map[string]string{}
	// ConfigApps - List of saved discord app IDs
	ConfigApps = map[string]string{}
)

// InitConfigFile ...
func InitConfigFile() {
	// Set platform spesific config path
	if runtime.GOOS == "windows" {
		homeDir, err := os.UserHomeDir()
		if err != nil {
			log.Fatalln(Lang["errorGettingHomeFolder"], err)
		}
		ConfigDir = homeDir + "\\AppData\\Roaming\\" + AppName + "\\"
	} else {
		homeDir, err := os.UserHomeDir()
		if err != nil {
			log.Fatalln(Lang["errorGettingHomeFolder"], err)
		}
		ConfigDir = homeDir + "/.config/" + AppName + "/"
	}

	// Check Config
	configFileBuf := bytes.NewBuffer(nil)
	configFile, err := os.Open(ConfigDir + ConfigFileName)
	if err != nil {
		// Config file not found?
		if !os.IsExist(err) {
			Config = ConfigDefault
		} else {
			log.Fatalln(Lang["errorOpeningConfigFile"], err)
		}
	} else {
		// Get configs from file
		io.Copy(configFileBuf, configFile)
		json.Unmarshal(configFileBuf.Bytes(), &Config)
	}
	configFile.Close()

	// Check ConfigApps
	configAppsFileBuf := bytes.NewBuffer(nil)
	configAppsFile, err := os.Open(ConfigDir + ConfigAppsFileName)
	if err != nil {
		// Config file not found?
		if os.IsExist(err) {
			log.Fatalln(Lang["errorOpeningConfigFile"], err)
		}
	} else {
		// Get configs from file
		io.Copy(configAppsFileBuf, configAppsFile)
		json.Unmarshal(configAppsFileBuf.Bytes(), &ConfigApps)
	}
	configAppsFile.Close()

	// Check for missing entries
	for key := range ConfigDefault {
		if val, ok := Config[key]; ok {
			if len(val) < 1 {
				Config[key] = ConfigDefault[key]
			}
		} else {
			Config[key] = ConfigDefault[key]
		}
	}

	// Disable Debug
	if os.Getenv("Discord-RPC-Tray_DEBUG") != "true" && Config["debug"] == "off" {
		log.SetOutput(ioutil.Discard)
	}

	log.Println("Config: ", Config)
	log.Println("ConfigApps: ", ConfigApps)
	ConfigSave()
}

// SaveConfig - Saves configs to file
func ConfigSave() {
	log.Println(Lang["debugConfigSaving"])
	if err := os.MkdirAll(ConfigDir, os.ModePerm); err != nil {
		log.Fatalln(Lang["errorCreatingConfigFolder"], err)
	}

	// Config
	configFile, err := os.Create(ConfigDir + ConfigFileName)
	if err != nil {
		log.Fatalln(Lang["errorOpeningConfigFile"], err)
	}
	defer configFile.Close()
	configBytes, err := json.Marshal(Config)
	if err != nil {
		log.Fatalln(Lang["errorMarshalingConfigData"], err)
	}
	if _, err = configFile.Write(configBytes); err != nil {
		log.Fatalln(Lang["errorWritingConfigFile"], err)
	}
	if err = configFile.Sync(); err != nil {
		log.Fatalln(Lang["errorSyncingConfigFile"], err)
	}

	// ConfigApps
	configAppsFile, err := os.Create(ConfigDir + ConfigAppsFileName)
	if err != nil {
		log.Fatalln(Lang["errorOpeningConfigFile"], err)
	}
	defer configAppsFile.Close()
	configAppsBytes, err := json.Marshal(ConfigApps)
	if err != nil {
		log.Fatalln(Lang["errorMarshalingConfigData"], err)
	}
	if _, err = configAppsFile.Write(configAppsBytes); err != nil {
		log.Fatalln(Lang["errorWritingConfigFile"], err)
	}
	if err = configAppsFile.Sync(); err != nil {
		log.Fatalln(Lang["errorSyncingConfigFile"], err)
	}

	log.Println(Lang["debugConfigSaved"])
}
