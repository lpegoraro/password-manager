package password_generator

import "github.com/nu7hatch/gouuid"

type password_configuration struct {
	generationMethod string
	pattern string
	strengthFactor int8
}

func generate_password(password_configuration password_configuration) string {
	if password_configuration.generationMethod == "uuid" {
		return uuid(password_configuration.pattern)
	}
}


func uuid(pattern string) string {
	return pattern
}
