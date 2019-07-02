package main

import (
	"fmt"
	storage "github.com/lpegoraro/password-manager/storage"
)

func save(description string, username string, configuration PasswordConfiguration, password string) {
	storageType := configuration.Storage
	passwordStore := storage.PasswordEntry{
		Tag:      description,
		Username: username,
		Password: password,
	}
	switch storageType {
	case "NOT_ENCRYPTED_FILE":
		storage.SaveToFile(passwordStore)
		break
	case "OUTPUT":
		configuration.Output = true
		break
	}
	if configuration.Output {
		fmt.Println("Password Generated: " + password)
	}
}
