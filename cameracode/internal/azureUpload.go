package internal

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path"
	"strings"
	"time"
)

func UploadFile(fileNameWithPath string, azureConfig AzureBlobUploadAccess) error {

	info, err := os.Stat(fileNameWithPath)
	if err != nil {
		log.Printf("failed to get file size: %v: %v", fileNameWithPath, err)
		return err
	}
	fileSize := info.Size()

	client := &http.Client{}
	// sas token from azure portal - beware it runs out over time
	fileName := path.Base(fileNameWithPath)

	// ignore leading question mark
	azureConfig.SasToken = strings.TrimPrefix(azureConfig.SasToken, "?")

	url := fmt.Sprintf("https://%v.blob.core.windows.net/%v/%v?%v",
		azureConfig.AccountName, azureConfig.ContainerName, fileName, azureConfig.SasToken)

	fileData, err := os.Open(fileNameWithPath)
	defer closeFile(fileData)

	if err != nil {
		log.Printf("failed to read file: %v: %v", fileNameWithPath, err)
		return err
	}

	req, err := http.NewRequest(http.MethodPut, url, fileData)

	req.Header.Set("Content-Type", "application/octet-stream")
	req.Header.Set("x-ms-blob-type", "BlockBlob")

	req.Header.Set("x-ms-version", "2015-02-21")
	req.Header.Set("x-ms-date", time.Now().UTC().Format(time.RFC1123))

	// content length is one automatically when req knows content size
	req.ContentLength = fileSize

	if err != nil {
		log.Printf("failed to build request: %v", err)
		return err
	}

	resp, err := client.Do(req)
	if err != nil {
		log.Printf("failed to execute request: %v", err)
		return err
	}

	defer closeResponse(resp.Body)

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("failed to read response: %v", err)
		return err
	}
	log.Printf("upload status: %v\n", resp.StatusCode)
	bodyText := string(body)
	if len(bodyText) > 0 {
		log.Printf("body returned: %v", bodyText)
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return fmt.Errorf("http status indicates error: %v", resp.StatusCode)
	}
	return nil
}

func closeResponse(body io.ReadCloser) {
	err := body.Close()
	if err != nil {
		log.Printf("failed to close body: %v", err)
	}
}

func closeFile(file *os.File) {
	err := file.Close()
	if err != nil {
		// log.Printf("failed to close file %v %v\n", file.Name(), err)
	}
}
