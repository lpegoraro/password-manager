package storage

import "log"

type PasswordEntry struct {
	Tag      string `json:"tag"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type StorePasswords struct {
	StoredPasswords map[string]PasswordEntry `json:"storedPasswords"`
}

type StorageStrategy interface {
	StorageSave(passwordEntry PasswordEntry, output bool)
	StorageGet(tag string, username string, output bool) string
}

func BuildStorage(storageType string) StorageStrategy {
	switch storageType {
	case "NOT_ENCRYPTED_FILE":
		return NotEncryptedFileStorageStrategy{}
	case "IMMUDB":
		immudb, err := NewImmuDbStorageStrategy()
		if err != nil {
			log.Fatalf("Error connecting to ImmuDB, %e", err)

		}
		return immudb
	case "":
		return NoSaveStrategy{}
	default:
		panic("Storage Type not found")
	}
}
