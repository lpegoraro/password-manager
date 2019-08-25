package main

import (
	"fmt"

	"github.com/lpegoraro/password-manager/storage"
)

func getStrategy(storageType string) StorageStrategy {
	dfs := DefaultStorageFactory{}
	return dfs.BuildStorage(storageType)
}

func save(description string, username string, configuration PasswordConfiguration, password string) {
	storageType := configuration.Storage
	passwordStore := storage.PasswordEntry{
		Tag:      description,
		Username: username,
		Password: password,
	}
	getStrategy(storageType).Save(passwordStore, configuration.Output)
	if configuration.Output {
		fmt.Println("Password Generated: " + password)
	}
}

func get(description string, username string, configuration PasswordConfiguration) string {
	storageType := configuration.Storage
	return getStrategy(storageType).Get(description, username, configuration.Output)

}
