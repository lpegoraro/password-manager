package password_generator

type password_configuration struct {
	generationMethod string
	pattern string
	strengthFactor int8
}

func generate_password(password_configuration password_configuration) string {
	if password_configuration.generationMethod == "uuid" {
		return uuid(password_configuration.strengthFactor, password_configuration.pattern)
	}
}


func uuid(strengthFactor int8, pattern string) string {
	if strengthFactor == 4 {
		return uuid.newV4()
	} else if strengthFactor == 5 {
		return uuid.newV5(pattern)
	}
}
