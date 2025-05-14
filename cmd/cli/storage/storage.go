package storage

import (
	"encoding/json"
	"os"
)

type Storage struct {
	Keys  string
	Codes []CodesEntry
}

type CodesEntry struct {
	Name string
	// Code usage per week
	Frequency int
	Length    int
	Code      string
}

func Load(codesFileAbs string) (Storage, error) {
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

func Save(storage *Storage, codesFileAbs string) error {
	theJson, err := json.MarshalIndent(storage, "", "  ")
	if err != nil {
		return err
	}

	err = os.WriteFile(codesFileAbs, theJson, 0644)
	if err != nil {
		return err
	}
	
	return nil
}
