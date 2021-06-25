package main

import (
	"log"
	"os"
)

func deleteFile(pathToFile string) {
	if _, err := os.Stat(pathToFile); os.IsNotExist(err) {
		return
	}
	err := os.Remove(pathToFile)
	if err != nil {
		log.Fatal(err)
	}
}

func fileExists(filepath string) bool {
	_, err := os.Stat(filepath)
	if os.IsNotExist(err) {
		return false
	}
	return true
}
