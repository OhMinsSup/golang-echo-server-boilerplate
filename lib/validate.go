package lib

import validator "gopkg.in/go-playground/validator.v9"

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) validate(i interface{}) error {
	return cv.validator.Struct(i)
}
