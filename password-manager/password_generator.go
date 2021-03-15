package main

import (
	"bufio"
	"io"
	"log"
	"os"
	"strings"
	"time"

	uuid "github.com/nu7hatch/gouuid"
)

type PasswordConfiguration struct {
	Method  string `json:"method"`
	Length  int32  `json:"password_length"`
	Seed    string `json:"seed"`
	Factor  int32  `json:"factor"`
	Storage string `json:"storage"`
	Output  bool   `json:"show_output"`
}

func GeneratePassword(passwordConfiguration PasswordConfiguration) string {
	switch passwordConfiguration.Method {
	case "uuid":
		return getUuid(passwordConfiguration.Factor, passwordConfiguration.Seed)
	case "hash":
		return getHash(passwordConfiguration)
	case "custom":
		return readPassword(os.Stdin)
	default:
		return ""
	}
}

func readPassword(in io.Reader) string {
	reader := bufio.NewReader(in)
	log.Println("Insert the password you want:")
	password, _ := reader.ReadString('\n')
	password = strings.Replace(password, "\n", "", -1)
	return password
}

func getHash(c PasswordConfiguration) string {
	pg := PasswordGeneration{
		config: config{
			flags:  All,
			length: c.Length,
		},
	}
	gen, err := pg.genPass()
	if err != nil {
		log.Fatal("Could not generate password", err)
	}
	return gen
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
