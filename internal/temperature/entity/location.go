package entity

import "regexp"

type Location struct {
	Cep        string
	Localidade string
}

func CEPValidation(cep string) error {
	re := regexp.MustCompile(`^\d{5}-\d{3}$`)
	if !re.MatchString(cep) {
		return ErrCEPNotValid
	}

	return nil
}
