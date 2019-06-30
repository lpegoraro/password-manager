package main

import (
	"fmt"
)

func save(description string, username string, configuration PasswordConfiguration, password string) {
	storageType := configuration.Storage
	passwordStore := PasswordEntry{
		Tag:      description,
		Username: username,
		Password: password,
	}
	switch storageType {
	case "NOT_ENCRYPTED_FILE":
		storage.saveToFile(passwordStore, configuration.)
		break
	case "OUTPUT":
		configuration.Output = true
		break
	}
	if configuration.Output {
		fmt.Println("Password Generated: " + password)
	}
}
