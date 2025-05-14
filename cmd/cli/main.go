package main

import (
	"log"
	"path/filepath"
	"physical-pin-code-generator/cmd/cli/generator"
	"physical-pin-code-generator/cmd/cli/histogram"
	"physical-pin-code-generator/cmd/cli/storage"
	"runtime"
	"strings"
)

const codesFile = "../../codes.json"

func main() {
	log.Println("Hello my friend")
	var codesFileAbs = getSourceDir() + "/" + codesFile

	log.Printf("I will use the file %s for codes\n", codesFileAbs)

	payload, err := storage.Load(codesFileAbs)
	if err != nil {
		panic(err)
	}

	log.Printf("Allowed keys: %s\n", payload.Keys)

	h := getHistogram(&payload)
	h.Print()

	generator.GenerateMissingCodes(&payload, h)

	h.Print()

	err = storage.Save(&payload, codesFileAbs)
	if err != nil {
		panic(err)
	}

	log.Println("All done")
}

func getSourceDir() string {
	_, file, _, _ := runtime.Caller(0)
	dir := filepath.Dir(file)
	return dir
}

func getHistogram(storage *storage.Storage) *histogram.Histogram {
	h := histogram.New()

	// add a bin for each allowed key
	allowedKeys := strings.Split(storage.Keys, "")
	for i := range allowedKeys {
		h.Add(allowedKeys[i])
	}

	// add the key usage counts
	for _, code := range storage.Codes {
		keys := strings.Split(code.Code, "")
		for i := range keys {
			// scale the usage by frequency
			h.AddInc(keys[i], code.Frequency)
		}
	}
	return h
}
