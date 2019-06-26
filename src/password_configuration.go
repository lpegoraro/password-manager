package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strings"
)

var COMMENDATION = "It is highly recommended to change the default by creating a config file, please run password_manager config"

func GetCurrentConfiguration() PasswordConfiguration {
	fileConfiguration := LoadFromFile()
	if fileConfiguration.generationMethod == "" {
		fmt.Println(COMMENDATION)
		return GetDefaultConfig()
	}
	return fileConfiguration
}

func GetDefaultConfig() PasswordConfiguration {
	return PasswordConfiguration{
		generationMethod: "uuid",
		seed:             "pwd_manager_test",
		strengthFactor:   5,
	}

}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func LoadFromFile() PasswordConfiguration {
	configFile := FindFile("./config/", strings.Split("password_config.json,", ","))
	configuration := PasswordConfiguration{}
	_ = json.Unmarshal([]byte(configFile), &configuration)
	fmt.Println(configuration.generationMethod)
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
