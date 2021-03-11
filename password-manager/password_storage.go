package main

import (
	"log"

	"github.com/lpegoraro/password-manager/storage"
)

func save(description string, username string, configuration PasswordConfiguration, password string) {
	storageType := configuration.Storage
	passwordStore := storage.PasswordEntry{
		Tag:      description,
		Username: username,
		Password: password,
	}
	storageStrategy := storage.BuildStorage(storageType)
	storageStrategy.StorageSave(passwordStore, configuration.Output)
	if configuration.Output {
		log.Println("Password Generated: " + password)
	}
}

func get(description string, username string, configuration PasswordConfiguration) string {
	storageType := configuration.Storage
	storageStrategy := storage.BuildStorage(storageType)
	return storageStrategy.StorageGet(description, username, configuration.Output)
}
