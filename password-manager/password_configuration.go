package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

var COMMENDATION = "It is highly recommended to change the default by creating a config file, please run password_manager config"
var homeDir = os.Getenv("HOME")
var CONFIG_FILE = homeDir + "/.secure/config/password_configuration.json"

func GetCurrentConfiguration() PasswordConfiguration {
	return LoadFromFile()
}

func GetDefaultConfig() PasswordConfiguration {
	return PasswordConfiguration{
		Method:  "uuid",
		Seed:    "pwd_manager_test",
		Factor:  4,
		Storage: "output",
		Output:  true,
	}
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func CreateConfigFile(method string, seed string, factor int32, storageType string) {
	configuration := PasswordConfiguration{
		Method:  method,
		Seed:    seed,
		Factor:  factor,
		Storage: storageType,
		Output:  false,
	}
	configurationJson, err := json.Marshal(configuration)
	fmt.Println(configurationJson)
	if err != nil {
		fmt.Println("Error on marshalling json")
		panic(err)
	}
	errorOnWriting := ioutil.WriteFile(CONFIG_FILE, configurationJson, 0644)
	fmt.Print("wrote file")
	if errorOnWriting != nil {
		fmt.Println("Error on saving the file")
		panic(errorOnWriting)
	}
}

func LoadFromFile() PasswordConfiguration {
	configFile, err := ioutil.ReadFile(CONFIG_FILE)
	if err != nil {
		fmt.Println(COMMENDATION)
		fmt.Println(err)
		return GetDefaultConfig()
	}
	configuration := PasswordConfiguration{}
	err2 := json.Unmarshal(configFile, &configuration)
	check(err2)
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
	return file
}
