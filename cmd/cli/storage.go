package main

import (
	"encoding/json"
	"os"
)

type Storage struct {
	Keys  string
	Codes []CodesEntry
}

type CodesEntry struct {
	Name      string
	Frequency float32
	Length    int
	Code      string
}

func load(codesFileAbs string) (Storage, error) {
	theJson, err := os.ReadFile(codesFileAbs)
	if err != nil {
		return Storage{}, err
	}

	var payload Storage
	err = json.Unmarshal(theJson, &payload)
	if err != nil {
		return Storage{}, err
	}

	return payload, nil
}
