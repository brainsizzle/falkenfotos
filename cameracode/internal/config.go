package internal

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

type Config struct {
	IntervalSeconds       int                   `json:"intervalSeconds"`
	ImageFileName         string                `json:"imageFileName"`
	AzureBlobUploadAccess AzureBlobUploadAccess `json:"azureBlobUploadAccess"`
}

type AzureBlobUploadAccess struct {
	AccountName   string `json:"accountName"`
	SasToken      string `json:"sasToken"`
	ContainerName string `json:"containerName"`

	// sasToken needes:
	// allowed Services: Blob
	// allowed resource type: Object
	// allowed permissions: Write, Add, Create
	// https only
}

func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	if err != nil {
		return false
	}
	return !info.IsDir()
}

func ReadConfig(configFilePath string) *Config {

	absolutePath, err := filepath.Abs(configFilePath)
	if err != nil {
		log.Printf("file path not found: %v", err)
		return nil
	}

	if !fileExists(absolutePath) {
		log.Printf("no config found in %v", absolutePath)
		return nil
	}

	fileData, err := ioutil.ReadFile(absolutePath)
	if err != nil {
		log.Printf("could not read file: %v ,%v", absolutePath, err)
		return nil
	}
	falcomCamConfig := Config{}
	err = json.Unmarshal(fileData, &falcomCamConfig)
	if err != nil {
		log.Printf("failed to parse config file: %v ,%v", absolutePath, err)
		return nil
	}
	log.Printf("config found in %v\n", absolutePath)
	return &falcomCamConfig
}
