package request

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

const (
	passwordMinLength = 8
)

type Register struct {
	Email    string `json:"email"    example:"test@example.com"`
	Password string `json:"password" example:"s3cr3tpw"`
}

func (r Register) Validate() error {
	return validation.ValidateStruct(&r,
		validation.Field(&r.Email, validation.Required, is.Email),
		validation.Field(&r.Password, validation.Required, validation.Length(passwordMinLength, 0)),
	)
}

type Login struct {
	Email    string `json:"email"    example:"test@example.com"`
	Password string `json:"password" example:"s3cr3tpw"`
}

func (r Login) Validate() error {
	return validation.ValidateStruct(&r,
		validation.Field(&r.Email, validation.Required),
		validation.Field(&r.Password, validation.Required),
	)
}
