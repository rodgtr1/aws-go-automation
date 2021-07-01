package utils

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func DeleteFile(pathToFile string) {
	if _, err := os.Stat(pathToFile); os.IsNotExist(err) {
		return
	}
	err := os.Remove(pathToFile)
	if err != nil {
		log.Fatal(err)
	}
}

func FileExists(filepath string) bool {
	_, err := os.Stat(filepath)
	if os.IsNotExist(err) {
		return false
	}
	return true
}

func ReplaceWordInFile(file string, wordToReplace string, replacementText string) {
	input, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	output := bytes.Replace(input, []byte(wordToReplace), []byte(replacementText), -1)

	if err = ioutil.WriteFile(".env", output, 0666); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
