package main

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"os"
)

var (
	// AppConfigFileName ...
	AppConfigFileName string = "config.json"
	// AppConfigFileDir ...
	AppConfigFileDir string = GetConfigPath()
	// AppConfigFilePath ...
	AppConfigFilePath string = AppConfigFileDir + AppConfigFileName

	// ConfigDefault ...
	ConfigDefault map[string]string = map[string]string{
		"lang": "en",
	}
	// Config - Settings of the app
	Config map[string]string = map[string]string{}
	// ConfigApps - List of saved discord app IDs
	ConfigApps map[string]string = map[string]string{}
)

// InitConfigFile ...
func InitConfigFile() {
	configFile, err := os.Open(AppConfigFilePath)
	defer configFile.Close()
	configFileBuf := bytes.NewBuffer(nil)
	if err != nil {
		// Config file not found
		if os.IsExist(err) != true {
			// Writing new config file
			log.Println("Config file not found. Creating new config file...")
			err := os.MkdirAll(AppConfigFileDir, os.ModePerm)
			if err != nil {
				log.Fatalln("Error on creating config folder: ", err)
			}
			configFile.Close()
			configFile, err = os.Create(AppConfigFilePath)
			if err != nil {
				log.Fatalln("Error on creating config file: ", err)
			}
			configBytes, err := json.Marshal(ConfigDefault)
			if err != nil {
				log.Fatalln("Error on marshaling config data: ", err)
			}
			_, err = configFile.Write(configBytes)
			if err != nil {
				log.Fatalln("Error on writing config file: ", err)
			}
			err = configFile.Sync()
			if err != nil {
				log.Fatalln("Error on syncing config file: ", err)
			}
			log.Println("New config file was created.")
		} else {
			log.Fatalln("Error on opening config file: ", err)
		}
	}

	// Get configs from file
	io.Copy(configFileBuf, configFile)
	json.Unmarshal(configFileBuf.Bytes(), &Config)
	configFile.Close()

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
	log.Println(Config)
	SaveConfig()
}

// SaveConfig - Saves configs to file
func SaveConfig() {
	log.Println("Saving config file...")
	configFile, err := os.Create(AppConfigFilePath)
	if err != nil {
		log.Fatalln("Error on opening config file: ", err)
	}
	defer configFile.Close()
	log.Println(Config)
	configBytes, err := json.Marshal(Config)
	if err != nil {
		log.Fatalln("Error on marshaling config data: ", err)
	}
	_, err = configFile.Write(configBytes)
	if err != nil {
		log.Fatalln("Error on writing config file: ", err)
	}
	err = configFile.Sync()
	if err != nil {
		log.Fatalln("Error on syncing config file: ", err)
	}
	log.Println("Config file was saved.")
}
