package main

import (
	"os"
	"testing"

	"github.com/lpegoraro/password-manager/storage"
	"github.com/stretchr/testify/assert"
)

// TestMain setup testing environment
func TestMain(m *testing.M) {
	// override config file path
	CONFIG_FILE = "./test-config.json"
	storage.STORAGE_FILE = "./test-passwordmanager.json"
	result := m.Run()
	removeTmpFile()
	os.Exit(result)
}

func removeTmpFile() {
	os.Remove(CONFIG_FILE)
	os.Remove(storage.STORAGE_FILE)
}

// TestParseArgsFail pass invalid argument to parseArgs function
func TestParseArgsNoArg(t *testing.T) {
	invalidArg := []string{}
	err := parseArgs(invalidArg)
	if assert.Error(t, err) {
		assert.Contains(t, err.Error(), "No valid argument found")
	}
}

// // TestParseArgsFail pass invalid argument to parseArgs function
func TestParseArgsFail(t *testing.T) {
	invalidArg := []string{"-x"}
	err := parseArgs(invalidArg)
	if assert.Error(t, err) {
		assert.Contains(t, err.Error(), "No valid argument found")
	}
}

// TestParsArgsMultiple valid argument passed to function and config file
// will be created
func TestParseArgsMultiple(t *testing.T) {
	argConfig := []string{"-c", "uuid", "'passphrase'", "4", "NOT_ENCRYPTED_FILE", "-g"}
	if assert.NoError(t, parseArgs(argConfig)) {
		// create configuration file
		assert.FileExists(t, CONFIG_FILE)
	}
}

// TestParsArgsConfig valid argument passed to function and config file
// will be created
func TestParseArgsConfig(t *testing.T) {
	argConfig := []string{"-c", "uuid", "'passphrase'", "4", "NOT_ENCRYPTED_FILE"}
	if assert.NoError(t, parseArgs(argConfig)) {
		// create configuration file
		if assert.FileExists(t, CONFIG_FILE) {
			// TODO check if there is correct configuration
		}
	}
}

func TestParseArgsHelp(t *testing.T) {
	argHelp := []string{"--help"}
	assert.NoError(t, parseArgs(argHelp))
	// TODO parse help text from stdin
}

func TestParseArgsHandleAdd(t *testing.T) {
	argConfig := []string{"-c", "uuid", "'passphrase'", "4", "NOT_ENCRYPTED_FILE"}
	if assert.NoError(t, parseArgs(argConfig)) {
		// create configuration file
		if assert.FileExists(t, CONFIG_FILE) {
			// TODO check if there is correct configuration
		}
	}
	argAdd := []string{"-a", "this", "newPhrase"}
	if assert.NoError(t, parseArgs(argAdd)) {

	}
}

func TestParseArgsHandleGet(t *testing.T) {
	argConfig := []string{"-c", "uuid", "'passphrase'", "4", "NOT_ENCRYPTED_FILE"}
	if assert.NoError(t, parseArgs(argConfig)) {
		// create configuration file
		if assert.FileExists(t, CONFIG_FILE) {
			// TODO check if there is correct configuration
		}
	}
	argAdd := []string{"-a", "this", "newPhrase"}
	if assert.NoError(t, parseArgs(argAdd)) {

	}
	argGet := []string{"-g", "this", "newPhrase"}
	if assert.NoError(t, parseArgs(argGet)) {

	}
}
