package main

import (
	"bufio"
	"fmt"
	"github.com/nu7hatch/gouuid"
	"os"
	"strings"
)

type PasswordConfiguration struct {
	Method string `json: "method"`
	Seed   string `json: "seed"`
	Factor int8   `json: "factor"`
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
	default :
		return ""
	}
}

func getCert(passwordConfiguration PasswordConfiguration) string {
	return ""
}

func getUuid(Factor int8, pattern string) string {
	if Factor == 4 {
		u4, err := uuid.NewV4()
		if err != nil {
			return ""
		}
		return u4.String()
	} else if Factor == 5 {
		u5, err := uuid.NewV5(uuid.NamespaceURL, []byte(pattern))
		if err != nil {
			return ""
		}
		return u5.String()
	}
	return ""
}
