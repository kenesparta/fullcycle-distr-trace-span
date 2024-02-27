package entity

import "regexp"

type Location struct {
	Cep        string
	Localidade string
}

func CEPValidation(cep string) error {
	re := regexp.MustCompile(`^\d{8}$`)
	if !re.MatchString(cep) {
		return ErrCEPNotValid
	}

	return nil
}
