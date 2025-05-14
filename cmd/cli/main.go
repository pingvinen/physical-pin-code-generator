package main

import (
	"log"
	"path/filepath"
	"runtime"
)

const codesFile = "../../codes.json"

func main() {
	log.Println("Hello my friend")
	var codesFileAbs = getSourceDir() + "/" + codesFile

	log.Printf("I will use the file %s for codes\n", codesFileAbs)

	payload, err := load(codesFileAbs)
	if err != nil {
		panic(err)
	}

	log.Printf("Allowed keys: %s\n", payload.Keys)

	log.Println("All done")
}

func getSourceDir() string {
	_, file, _, _ := runtime.Caller(0)
	dir := filepath.Dir(file)
	return dir
}
