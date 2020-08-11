package main

import (
	"regexp"
	"strings"
	"testing"
)

func TestGeneratePassword(t *testing.T) {

	t.Run("returns empty string on unkown PassWordConfiguration method", func(t *testing.T) {
		got := GeneratePassword(PasswordConfiguration{
			Method: "",
		})
		if got != "" {
			t.Errorf("expected '' but got '%v'", got)
		}
	})

	assertCorrectUUID := func(t *testing.T, uuid string) {
		t.Helper()
		matched, err := regexp.MatchString(`\b[0-9a-f]{8}\b-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-\b[0-9a-f]{12}\b`, uuid)
		if err != nil {
			t.Errorf("error matching regex for value %v", uuid)
		}
		if !matched {
			t.Errorf("expected a valid uuid but got '%v'", uuid)
		}
	}

	t.Run("returns valid UUID v4 on uuid PassWordConfiguration method", func(t *testing.T) {
		got := GeneratePassword(PasswordConfiguration{
			Method: "uuid",
			Size:   10,
			Seed:   "123",
		})
		assertCorrectUUID(t, got)
	})

	t.Run("returns valid UUID v5 on uuid PassWordConfiguration method", func(t *testing.T) {
		got := GeneratePassword(PasswordConfiguration{
			Method: "uuid",
			Size:   10,
			Seed:   "123",
		})
		assertCorrectUUID(t, got)
	})

	t.Run("returns different UUID v4 on uuid PassWordConfiguration method", func(t *testing.T) {
		got := GeneratePassword(PasswordConfiguration{
			Method: "uuid",
			Size:   10,
			Seed:   "123",
		})
		got2 := GeneratePassword(PasswordConfiguration{
			Method: "uuid",
			Size:   10,
			Seed:   "123",
		})
		if got == got2 {
			t.Errorf("expected different values for multiple call but received duplicates '%v'", got)
		}
	})

	t.Run("returns different UUID v5 on uuid PassWordConfiguration method", func(t *testing.T) {
		got := GeneratePassword(PasswordConfiguration{
			Method: "uuid",
			Size:   10,
			Seed:   "123",
		})
		got2 := GeneratePassword(PasswordConfiguration{
			Method: "uuid",
			Size:   10,
			Seed:   "123",
		})
		if got == got2 {
			t.Errorf("expected different values for multiple call but received duplicates '%v'", got)
		}
	})

	t.Run("readPassword returns entered password", func(t *testing.T) {
		input := "my_awesome_password"
		got := readPassword(strings.NewReader(input + "\n"))
		if input != got {
			t.Errorf("expected '%v' but got '%v'", input, got)
		}
	})

}
