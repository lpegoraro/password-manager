package storage

import "fmt"

type NoSaveStrategy struct {
}

func (nss NoSaveStrategy) StorageSave(passwordEntry PasswordEntry, output bool) {
	log.Println("Mock Saving the password")
	log.Println(passwordEntry.Tag)
	log.Println(passwordEntry.Username)
	log.Println(passwordEntry.Password)
}

func (nss NoSaveStrategy) StorageGet(tag string, username string, output bool) string {
	log.Println("Mock Getting the password")
	log.Println("The mocked password will return")
	return "MockPassword!23$"
}
