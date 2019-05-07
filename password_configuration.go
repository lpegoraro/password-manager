package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
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
		strengthFactor:   6,
	}

}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func LoadFromFile() PasswordConfiguration {
	configFile, err := ioutil.ReadFile("./config/*.json")
	check(err)
	// TODO Remove this log
	fmt.Println(string(configFile))
	configuration := PasswordConfiguration{}
	_ = json.Unmarshal([]byte(configFile), &configuration)
	// TODO Remove this log
	fmt.Println("Configuration loaded: " + configuration.generationMethod + ", " +
		configuration.seed)
	return configuration
}
