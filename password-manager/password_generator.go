package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
	"time"

	uuid "github.com/nu7hatch/gouuid"
)

type PasswordConfiguration struct {
	Method  string `json: "method,omitempty"`
	Seed    string `json: "seed,omitempty"`
	Factor  int32  `json: "factor,omitempty"`
	Storage string `json: "storage,omitempty"`
	Output  bool   `json: "show_output,omitempty"`
}

func GeneratePassword(passwordConfiguration PasswordConfiguration) string {
	switch passwordConfiguration.Method {
	case "uuid":
		return getUuid(passwordConfiguration.Factor, passwordConfiguration.Seed)
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

func getUuid(Factor int32, pattern string) string {
	entropy := pattern + time.Now().String()
	if Factor == 4 {
		u4, err := uuid.NewV4()
		if err != nil {
			return ""
		}
		return u4.String()
	} else if Factor == 5 {
		u5, err := uuid.NewV5(uuid.NamespaceURL, []byte(entropy))
		if err != nil {
			return ""
		}
		return u5.String()
	}
	return ""
}
