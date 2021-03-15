package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
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
		Length:  15,
		Factor:  4,
		Storage: "output",
		Output:  true,
	}
}

func errorCheck(e error) {
	if e != nil {
		log.Fatal("Error found", e)
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
	log.Println(configurationJson)
	if err != nil {
		log.Fatal("Error on marshalling json", err)
	}
	errorOnWriting := ioutil.WriteFile(CONFIG_FILE, configurationJson, 0644)
	fmt.Print("wrote file")
	if errorOnWriting != nil {
		log.Fatal("Error on saving the file", errorOnWriting)
	}
}

func LoadFromFile() PasswordConfiguration {
	configFile, err := ioutil.ReadFile(CONFIG_FILE)
	if err != nil {
		log.Println(COMMENDATION)
		log.Println(err)
		return GetDefaultConfig()
	}
	configuration := PasswordConfiguration{}
	err = json.Unmarshal(configFile, &configuration)
	errorCheck(err)
	return configuration
}
