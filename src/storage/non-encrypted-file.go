package storage

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"path/filepath"
)

type PasswordEntry struct {
	Tag      string `json:"tag"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type PasswordJsonStore struct {
	key string `json:"tag"`
	password PasswordEntry `json:"password"`
}

var STORAGE_FILE = "~/.secure/.passwordmanager.json"


func saveToFile(passwordEntry PasswordEntry, filePath string) {
	ioutil.
}

func findInFile(tag string) string {
	ioutil.ReadFile(STORAGE_FILE)
}
