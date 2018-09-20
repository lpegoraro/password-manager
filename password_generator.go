package password_manager

import "github.com/nu7hatch/gouuid"

type PasswordConfiguration struct {
	generationMethod string
	seed             string
	strengthFactor   int8
}

func GeneratePassword(passwordConfiguration PasswordConfiguration) string {
	if passwordConfiguration.generationMethod == "uuid" {
		return getUuid(passwordConfiguration.strengthFactor, passwordConfiguration.seed)
	} else if passwordConfiguration.generationMethod == "cert" {
		return getCert(passwordConfiguration)
	} else {
		return ""
	}
}

func getCert(passwordConfiguration PasswordConfiguration) string {
	return ""
}

func getUuid(strengthFactor int8, pattern string) string {
	if strengthFactor == 4 {
		u4, err := uuid.NewV4()
		if err != nil {
			return ""
		}
		return u4.String()
	} else if strengthFactor == 5 {
		u5, err := uuid.NewV5(uuid.NamespaceURL, []byte(pattern))
		if err != nil {
			return ""
		}
		return u5.String()
	}
}
