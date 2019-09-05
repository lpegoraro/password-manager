package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	uuid "github.com/nu7hatch/gouuid"
)

type PasswordConfiguration struct {
	Method  string `json: "method"`
	Seed    string `json: "seed"`
	Factor  int32  `json: "factor"`
	Storage string `json: "storage"`
	Output  bool   `json: "show_output"`
}

func GeneratePassword(passwordConfiguration PasswordConfiguration) string {
	switch passwordConfiguration.Method {
	case "uuid":
		return getUuid(passwordConfiguration.Factor, passwordConfiguration.Seed)
	case "cert":
		return getCert(passwordConfiguration)
	case "custom":
		reader := bufio.NewReader(os.Stdin)
		fmt.Println("Insert the password you want:")
		password, _ := reader.ReadString('\n')
		password = strings.Replace(password, "\n", "", -1)
		return password
	default:
		return ""
	}
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
