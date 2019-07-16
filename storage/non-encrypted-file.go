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

const storageFile = os.Getenv("HOME") + "/.secure/.passwordmanager.json"

const NotEncryptedFileStorageStrategy = StorageStrategy{
	Save: func(passwordEntry PasswordEntry) {
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
		fmt.Println("Saved Passwords file " + string(jsonSavedPasswords))
		errorOnWriting := ioutil.WriteFile(storageFile, jsonSavedPasswords, 0644)
		if errorOnWriting != nil {
			fmt.Println("Error on saving the file")
			panic(errorOnWriting)
		}
	},
	Get: func(tag string, username string) string {
		file, err := ioutil.ReadFile(storageFile)
		if err != nil {
			fmt.Println("404 - Password not found!")
			return ""
		}
		savedPasswords := StorePasswords{}
		err2 := json.Unmarshal(file, &savedPasswords)
		if err2 != nil {
			panic(err2)
		}
		for tagKey, passwordValue := range savedPasswords.StoredPasswords {
			if tag == tagKey {
				return passwordValue.Password
			}
		}
		fmt.Println("404 - Password not found!")
		return ""
	},
}
