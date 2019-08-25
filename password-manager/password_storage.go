package main

import (
	"fmt"

	"github.com/lpegoraro/password-manager/storage"
)

type Storage interface {
	Save(description string, username string, password string)
	Get(description string, username string) string
}

type StorageFactory interface {
	GetStorage(storageType string) Storage
}

func save(description string, username string, configuration PasswordConfiguration, password string) {
	storageType := configuration.Storage
	passwordStore := storage.PasswordEntry{
		Tag:      description,
		Username: username,
		Password: password,
	}
	switch storageType {
	case "NOT_ENCRYPTED_FILE":
		storage.Save(passwordStore, configuration.Output)
		break
	case "OUTPUT":
		configuration.Output = true
		break
	}
	if configuration.Output {
		fmt.Println("Password Generated: " + password)
	}
}
