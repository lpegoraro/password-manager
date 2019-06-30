package storage

type PasswordEntry struct {
	Tag      string `json:"tag"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type PasswordJsonStore struct {
	key string `json:"tag"`
	password PasswordEntry `json:"password"`
}

var storageFile = "~/.secure/.passwordmanager.json"


func saveToFile(passwordEntry PasswordEntry, filePath string) {

}
