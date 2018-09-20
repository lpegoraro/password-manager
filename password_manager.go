package password_generator

import (
	"os"
	"fmt"
)

type key_secret struct {
	description, username  string
	tags []string
}

type config_argument struct {
	singleLetter, name, description string
}

var list_of_commands = "\n\thelp | -h: Prints this message\n" +
	"\tversion | -v: Print the version of the app\n" +
	"\tget | -g {DESCRIPTION} {USERNAME} {OPTIONS}: Copy the password to the clipboard, for more information use `password_manager get help\n" +
	"\tadd | -a {DESCRIPTION} {USERNAME} {OPTIONS}: Add a new password entry, for more information use `password_manager add help\n"+
	"\tconfig | -c {OPTIONS}: Configure encryption or password generation method\n"

var HELP_COMMAND = config_argument{
	singleLetter: "-h",
	name: "help",
	description: "Usage: `password_manager {COMMANDS} {OPTIONS}`" +
		"\n The command list is the below" +
		list_of_commands +
		"Hope this helps =)",
}

var ABOUT = config_argument{
	singleLetter: "-v",
	name: "version",
	description: "Password Manager in Go version 0.0.1",
}

func main() {
	arguments := os.Args[1:]
	parseArgs(arguments)
}

var GET_COMMAND = config_argument{
	singleLetter: "-g",
	name: "get",
	description: "Copy the password to the clipboard\n" +
		"Options:\n" +
		"\thelp: Print this help\n" +
		"\tverbose: Print the password in the console, not only copy to clipboard\n",
}

var ADD_COMMAND = config_argument{
	singleLetter: "-a",
	name: "add",
	description: "Add a new password entry\n" +
		"Options:\n" +
		"\thelp: Print this help\n" +
		"\tverbose: Print the password in the console after setting, not only copy to clipboard\n" +
		"\t",
}

var CONFIG_COMMAND = config_argument{

}

func parseArgs(arguments []string) {
	for index:=0; index < len(arguments); index++ {
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

func handleGet(arguments []string) {

}

func checkIfCommand(value string, command config_argument) bool {
	return value == command.singleLetter || value == "--"+command.name || value == command.name
}

