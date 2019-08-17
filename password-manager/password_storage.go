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
		storage.SaveToFile(passwordStore, configuration.Output)
		break
	case "OUTPUT":
		configuration.Output = true
		break
	}
	if configuration.Output {
		fmt.Println("Password Generated: " + password)
	}
}

func get(description string, username string, configuration PasswordConfiguration) string {
	storageType := configuration.Storage
	switch storageType {
	case "NOT_ENCRYPTED_FILE":
		return storage.FindInFile(description, username, configuration.Output)
	}
	return "not found"
}
