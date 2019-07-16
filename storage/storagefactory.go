package storage

type StorageStrategy interface {
	Save func(description string, username string, password string)
	Get func(description string, username string) string
}

func (st StorageFactory) GetStorage(storageType string) Storage {
	switch storageType {
	case "NON_ENCRYPTED_FILE":
		return NonEcryptedFileStorage
	case "":
		return NotSave
	default:
		panic("Storage Type not found")
	}
}

const NonEcryptedFileStorage = Storage{}

const NotSave = Storage{}
