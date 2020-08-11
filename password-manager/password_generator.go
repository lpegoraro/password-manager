package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"

	uuid "github.com/nu7hatch/gouuid"
)

type PasswordConfiguration struct {
	Method  string `json: "method"`
	Seed    string `json: "seed"`
	Size    int32  `json: "size"`
	Storage string `json: "storage"`
	Output  bool   `json: "show_output"`
}

func GeneratePassword(passwordConfiguration PasswordConfiguration) string {
	switch passwordConfiguration.Method {
	case "uuid":
		return getUuid()
	case "cert":
		return getCert(passwordConfiguration)
	case "custom":
		return readPassword(os.Stdin)
	default:
		return ""
	}
}

func readPassword(in io.Reader) string {
	reader := bufio.NewReader(in)
	fmt.Println("Insert the password you want:")
	password, _ := reader.ReadString('\n')
	password = strings.Replace(password, "\n", "", -1)
	return password
}

func getCert(passwordConfiguration PasswordConfiguration) string {
	return ""
}

func getUuid() string {
	u4, err := uuid.NewV4()
	if err != nil {
		return ""
	}
	return u4.String()
}
