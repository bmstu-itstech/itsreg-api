package model

import validation "github.com/go-ozzo/ozzo-validation/v4"

type Option struct {
	Text string `json:"text" example:"Опция 1"`
	Next int64  `json:"next" example:"1"`
}

func (o Option) Validate() error {
	return validation.ValidateStruct(&o,
		validation.Field(&o.Text, validation.Required),
		validation.Field(&o.Next, validation.Required),
	)
}
