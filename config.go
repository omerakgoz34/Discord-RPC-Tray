package main

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"os"
)

var (
	// Config ...
	Config map[string]string = map[string]string{
		"lang": "en",
	}
	// ConfigFile ...
	ConfigFile *os.File
	// AppConfigFileName ...
	AppConfigFileName string = "config.json"
	// AppConfigFileDir ...
	AppConfigFileDir string = GetConfigPath()
	// AppConfigFilePath ...
	AppConfigFilePath string = AppConfigFileDir + AppConfigFileName
)

// InitConfigFile ...
func InitConfigFile() {
	ConfigFile, err := os.Open(AppConfigFilePath)
	configFileBuf := bytes.NewBuffer(nil)
	if err != nil {
		if !os.IsExist(err) {
			log.Println("Config file not found. Creating new config file...")
			err := os.MkdirAll(AppConfigFileDir, os.ModePerm)
			if err != nil {
				log.Fatalln("Error on creating config folder: ", err)
			}
			ConfigFile, err := os.Create(AppConfigFilePath)
			if err != nil {
				log.Fatalln("Error on creating config file: ", err)
			}
			configBytes, err := json.Marshal(Config)
			if err != nil {
				log.Fatalln("Error on marshaling config data! ", err)
			}
			_, err = ConfigFile.Write(configBytes)
			if err != nil {
				log.Fatalln("Error on writing config file! ", err)
			}
			err = ConfigFile.Sync()
			if err != nil {
				log.Fatalln("Error on syncing config file! ", err)
			}
			log.Println("New config file was created.")
		} else {
			log.Fatalln("Error on opening config file! ", err)
		}
	}
	io.Copy(configFileBuf, ConfigFile)
	json.Unmarshal(configFileBuf.Bytes(), &Config)
	log.Println(Config)
}
