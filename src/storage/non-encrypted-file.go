package storage

type PasswordStore struct {
	Tag      string `json:"tag"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func saveToFile(passwordStore PasswordStore, filePath string) {

}
