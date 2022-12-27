package main

import (
	"cameracode/internal"
	"errors"
	"log"
	"time"
)

func main() {

	config := internal.ReadConfig("./falconcam.conf")
	if config == nil {
		panic(errors.New("config not found"))
	} else {
		log.Printf("config: %+v", *config)
	}

	for {
		err := internal.TakePicture(config.ImageFileName)
		if err != nil {
			log.Printf("take picture failed: %v", err)
			time.Sleep(time.Duration(config.IntervalSeconds) * time.Second)
			continue
		}

		err = internal.UploadFile(config.ImageFileName, config.AzureBlobUploadAccess)
		if err != nil {
			log.Printf("upload file failed: %v", err)
			time.Sleep(time.Duration(config.IntervalSeconds) * time.Second)
			continue
		}

		time.Sleep(time.Duration(config.IntervalSeconds) * time.Second)
	}
}
