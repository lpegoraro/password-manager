package storage

type PasswordEntry struct {
	Tag      string `json:"tag"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type StorePasswords struct {
	StoredPasswords map[string]PasswordEntry `json:"storedPasswords"`
}

type StorageStrategy interface {
	Save(passwordEntry PasswordEntry, output bool)
	Get(tag string, username string, output bool) string
}

func GetStorage(storageType string) StorageStrategy {
	switch storageType {
	case "NON_ENCRYPTED_FILE":
		return NotEncryptedFileStorageStrategy{}
	case "":
		return NoSaveStrategy{}
	default:
		panic("Storage Type not found")
	}
}
