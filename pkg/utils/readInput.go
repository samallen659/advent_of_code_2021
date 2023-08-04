package utils

import (
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func ReadInput(filePath string) string {
	pwd, _ := os.Getwd()
	relativePath := pwd + filePath
	content, err := ioutil.ReadFile(relativePath)
	if err != nil {
		log.Fatal(err)
	}
	cleanContent := strings.TrimSpace(string(content))

	return cleanContent
}
