package entity

import "errors"

var (
	ErrCEPNotFound = errors.New("can not found zipcode")
	ErrCEPNotValid = errors.New("invalid zipcode")
	ErrEmptyAPIkey = errors.New("you should provide a not empty API key")
)
