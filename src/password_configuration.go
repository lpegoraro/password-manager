package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"path/filepath"
)

var COMMENDATION = "It is highly recommended to change the default by creating a config file, please run password_manager config"

func GetCurrentConfiguration() PasswordConfiguration {
	fileConfiguration := LoadFromFile()
	if fileConfiguration.Method == "" {
		fmt.Println(COMMENDATION)
		return GetDefaultConfig()
	}
	return fileConfiguration
}

func GetDefaultConfig() PasswordConfiguration {
	return PasswordConfiguration {
		Method: "uuid",
		Seed:   "pwd_manager_test",
		Factor: 5,
	}
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func LoadFromFile() PasswordConfiguration {
	configFile, err := ioutil.ReadFile("./config/password_configuration.json")
	check(err)
	configuration := PasswordConfiguration{}
	err2 := json.Unmarshal(configFile, &configuration)
	check(err2)
	// TODO Remove this log
	fmt.Println("Configuration loaded: " + configuration.Method + ", " + configuration.Seed)
	return configuration
}

func FindFile(targetDir string, pattern []string) []byte {
	foundPath := ""
	for _, v := range pattern {
			matches, err := filepath.Glob(targetDir + v)
			check(err)
			if err != nil {
				fmt.Println(err)
			}

			if len(matches) != 0 {
				fmt.Println("Found : ", matches)
				foundPath = matches[0]
			}
	}
	file, err := ioutil.ReadFile(foundPath)
	check(err)
	return file;
}
