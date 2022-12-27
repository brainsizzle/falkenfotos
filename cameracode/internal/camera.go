package internal

import (
	"log"
	"os/exec"
	"strings"
)

func TakePicture(fileName string) error {
	prg := "./capture.py"
	arg1 := fileName

	cmd := exec.Command(prg, arg1)
	stdout, err := cmd.Output()

	if err != nil {
		log.Printf("failed to take picture %v", err)
		return err
	}
	outputText := string(stdout)
	if len(outputText) > 0 {
		outputText = strings.TrimSuffix(outputText, "\n")
		log.Printf("picture taken %v - output: %v ", fileName, outputText)
	} else {
		log.Printf("picture taken %v", fileName)
	}

	return nil
}
