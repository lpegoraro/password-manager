package main

import (
	"fmt"
	"storage"
)

func save(description string, username string, configuration PasswordConfiguration, password string) {
	storageType := configuration.Storage
	passwordStore := PasswordStore{
		Tag:      description,
		Username: username,
		Password: password,
	}
	switch storageType {
	case "NOT_ENCRYPTED_FILE":
		saveToFile(passwordStore, configuration.)
		break
	case "OUTPUT":
		configuration.Output = true
		break
	}
	if configuration.Output {
		fmt.Println("Password Generated: " + password)
	}
}
