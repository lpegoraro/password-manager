package storage

import "fmt"

type NoSaveStrategy struct {
}

func (nss NoSaveStrategy) StorageSave(passwordEntry PasswordEntry, output bool) {
	fmt.Println("Mock Saving the password")
	fmt.Println(passwordEntry.Tag)
	fmt.Println(passwordEntry.Username)
	fmt.Println(passwordEntry.Password)
}

func (nss NoSaveStrategy) StorageGet(tag string, username string, output bool) string {
	fmt.Println("Mock Getting the password")
	fmt.Println("The mocked password will return")
	return "MockPassword!23$"
}
