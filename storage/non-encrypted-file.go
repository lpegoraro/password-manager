package storage

import (
	"fmt"
	"encoding/json"
	"io/ioutil"
)

type PasswordEntry struct {
	Tag      string `json:"tag"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type StorePasswords struct {
	StoredPasswords []PasswordEntry `json:"storedPasswords"`
}

var STORAGE_FILE = "~/.secure/.passwordmanager.json"

func saveToFile(passwordEntry PasswordEntry) {
	file, err := ioutil.ReadFile(STORAGE_FILE)
	savedPasswords := StorePasswords{}
	if err != nil {
		savedPasswords.StoredPasswords = []PasswordEntry{}
	}
	err2 := json.Unmarshal(file, &savedPasswords)
	if err2 != nil {
		panic(err2)
	}
	savedPasswords.StoredPasswords = append(savedPasswords.StoredPasswords, passwordEntry)
	jsonSavedPasswords, err := json.Marshal(savedPasswords)
	fmt.Println("Saved Passwords file " + string(jsonSavedPasswords))
	errorOnWriting := ioutil.WriteFile(STORAGE_FILE, jsonSavedPasswords, 0644)
	if errorOnWriting != nil {
		fmt.Println("Error on saving the file")
		panic(errorOnWriting)
	}
}

func findInFile(tag string) string {
	ioutil.ReadFile(STORAGE_FILE)
	return "";
}
