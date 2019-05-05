package main

import (
	"fmt"
	"os"
)

type KeySecret struct {
	description, username string
	tags                  []string
}

type ConfigArgument struct {
	singleLetter, name, description string
}

var ListOfCommands = "\n\thelp | -h: Prints this message\n" +
	"\tversion | -v: Print the version of the app\n" +
	"\tget | -g {DESCRIPTION} {USERNAME} {OPTIONS}: Copy the password to the clipboard, for more information use `password_manager get help\n" +
	"\tadd | -a {DESCRIPTION} {USERNAME} {OPTIONS}: Add a new password entry, for more information use `password_manager add help\n" +
	"\tconfig | -c {OPTIONS}: Configure encryption or password generation method\n"

var HELP_COMMAND = ConfigArgument{
	singleLetter: "-h",
	name:         "help",
	description: "Usage: `password_manager {COMMANDS} {OPTIONS}`" +
		"\n The command list is the below" +
		ListOfCommands,
}

var ABOUT = ConfigArgument{
	singleLetter: "-v",
	name:         "version",
	description:  "Password Manager in Go version 0.0.1",
}

func main() {
	arguments := os.Args[1:]
	parseArgs(arguments)
}

var GET_COMMAND = ConfigArgument{
	singleLetter: "-g",
	name:         "get",
	description: "Copy the password to the clipboard\n" +
		"Options:\n" +
		"\thelp: Print this help\n" +
		"\tverbose: Print the password in the console, not only copy to clipboard\n",
}

var ADD_COMMAND = ConfigArgument{
	singleLetter: "-a",
	name:         "add",
	description: "Add a new password entry\n" +
		"Options:\n" +
		"\thelp: Print this help\n" +
		"\tverbose: Print the password in the console after setting, not only copy to clipboard\n" +
		"\t",
}

var CONFIG_COMMAND = ConfigArgument{
	singleLetter: "-a",
	name:         "add",
	description: "Configure the password generation\n" +
		"Options:\n" +
		"\thelp: Print this help\n" +
		"\tuuid\n" +
		"\t",
}

func parseArgs(arguments []string) {
	for index := 0; index < len(arguments); index++ {
		value := arguments[index]
		if checkIfCommand(value, HELP_COMMAND) {
			fmt.Println(HELP_COMMAND.description)
		} else if checkIfCommand(value, ABOUT) {
			fmt.Println(ABOUT.description)
		} else if checkIfCommand(value, GET_COMMAND) {
			handleGet(arguments)
		} else if checkIfCommand(value, ADD_COMMAND) {
			handleAdd(arguments)
		} else if checkIfCommand(value, CONFIG_COMMAND) {
			handleConfig(arguments)
		}
	}
	if len(arguments) == 0 {
		fmt.Println(ABOUT.description)
		fmt.Println(HELP_COMMAND.description)

	}
}
func handleConfig(arguments []string) {

}

func handleAdd(arguments []string) {

}

func handleGet(arguments []string) string {
	description := arguments[1]
	username := arguments[2]
	savedPassword := GetPassword(description, username)
	if savedPassword == "" {
		savedPassword = GeneratePassword(GetConfiguration(description, username))
	}
	return savedPassword
}
func GetConfiguration(description string, username string) PasswordConfiguration {
	return PasswordConfiguration{
		generationMethod: "uuid",
		seed:             "lPegz_password_manager_in_go",
		strengthFactor:   4,
	}
}
func GetPassword(description, username string) string {
	return ""
}

func checkIfCommand(value string, command ConfigArgument) bool {
	return value == command.singleLetter || value == "--"+command.name || value == command.name
}
