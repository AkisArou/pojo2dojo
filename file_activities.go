package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

const baseFolder = "mocks"

func makePathWithFolderFile(folder string, file string) string {
	return folder + "/" + file
}

func generateDartFile(dartClass string, path string) {
	_, err := os.Stat(baseFolder)

	if os.IsNotExist(err) {
		errDir := os.MkdirAll(baseFolder, 0755)
		if errDir != nil {
			log.Fatal(err)
		}
	}

	err2 := ioutil.WriteFile(path, []byte(dartClass), 0644)

	if err2 != nil {
		log.Fatal(err)
	}

	fmt.Println("Dart class file generated successfully!")
}

func readJavaFile(path string) string {
	data, err := ioutil.ReadFile(path)

	if err != nil {
		log.Fatal(err)
	}

	return string(data)
}
