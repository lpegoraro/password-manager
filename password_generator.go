package main

import (
	"bufio"
	"fmt"
	"github.com/nu7hatch/gouuid"
	"os"
	"strings"
)

type PasswordConfiguration struct {
	generationMethod string `json: "generationMethod"`
	seed             string `json: "seed"`
	strengthFactor   int8   `json: "strenghtFactor"`
}

func GeneratePassword(passwordConfiguration PasswordConfiguration) string {
	if passwordConfiguration.generationMethod == "uuid" {
		return getUuid(passwordConfiguration.strengthFactor, passwordConfiguration.seed)
	} else if passwordConfiguration.generationMethod == "cert" {
		return getCert(passwordConfiguration)
	} else if passwordConfiguration.generationMethod == "custom" {
		reader := bufio.NewReader(os.Stdin)
		fmt.Println("Insert the password you want:")
		password, _ := reader.ReadString('\n')
		password = strings.Replace(password, "\n", "", -1)
		return password
	} else {
		return ""
	}
}

func getCert(passwordConfiguration PasswordConfiguration) string {
	return ""
}

func getUuid(strengthFactor int8, pattern string) string {
	if strengthFactor == 4 {
		u4, err := uuid.NewV4()
		if err != nil {
			return ""
		}
		return u4.String()
	} else if strengthFactor == 5 {
		u5, err := uuid.NewV5(uuid.NamespaceURL, []byte(pattern))
		if err != nil {
			return ""
		}
		return u5.String()
	}
	return ""
}
