package storage

import (
	"encoding/json"
	"io/ioutil"
)

type PasswordEntry struct {
	Tag      string `json:"tag"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type StorePasswords struct {
	storedPasswords []PasswordEntry
}

var STORAGE_FILE = "~/.secure/.passwordmanager.json"

func saveToFile(passwordEntry PasswordEntry) {
	file, err := ioutil.ReadFile(STORAGE_FILE)
	savedPasswords := []StorePasswords{ storedPasswords: passwordEntry }
	if err != nil {
		panic(err)
	}
	err2 := json.Unmarshal(file, &savedPasswords)
	if err2 != nil {
		panic(err2)
	}
	append(savedPasswords.storedPasswords, passwordEntry)
	data := []byte()
	err := ioutil.WriteFile(STORAGE_FILE, data, 0644)

	if err != nil {
		panic(err)
	}
}

func findInFile(tag string) string {
	ioutil.ReadFile(STORAGE_FILE)
}
