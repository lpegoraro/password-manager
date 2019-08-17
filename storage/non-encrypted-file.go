package storage

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type PasswordEntry struct {
	Tag      string `json:"tag"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type StorePasswords struct {
	StoredPasswords map[string]PasswordEntry `json:"storedPasswords"`
}

var STORAGE_FILE = os.Getenv("HOME") + "/.secure/.passwordmanager.json"

func SaveToFile(passwordEntry PasswordEntry, output bool) {
	file, err := ioutil.ReadFile(STORAGE_FILE)
	savedPasswords := StorePasswords{}
	if err != nil {
		savedPasswords.StoredPasswords = make(map[string]PasswordEntry)
	} else {
		err2 := json.Unmarshal(file, &savedPasswords)
		if err2 != nil {
			panic(err2)
		}
	}
	passwordKey := passwordEntry.Tag + passwordEntry.Username
	// append the new passwordEntry to savedPasswords
	savedPasswords.StoredPasswords[passwordKey] = passwordEntry
	jsonSavedPasswords, err := json.Marshal(savedPasswords)
	if output {
		fmt.Println("Saved Passwords file " + string(jsonSavedPasswords))
	}
	errorOnWriting := ioutil.WriteFile(STORAGE_FILE, jsonSavedPasswords, 0644)
	if errorOnWriting != nil {
		fmt.Println("Error on saving the file")
		panic(errorOnWriting)
	}
}

func FindInFile(tag string, username string, output bool) string {
	file, err := ioutil.ReadFile(STORAGE_FILE)
	if err != nil {
		fmt.Println("404 - File of Password not found!")
		return ""
	}
	savedPasswords := StorePasswords{}
	err2 := json.Unmarshal(file, &savedPasswords)
	if err2 != nil {
		panic(err2)
	}
	for tagKey, passwordValue := range savedPasswords.StoredPasswords {
		if username+tag == tagKey {
			return passwordValue.Password
		}
	}
	fmt.Println("404 - Password not found!")
	return ""
}
