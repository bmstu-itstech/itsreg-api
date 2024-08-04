package model

import validation "github.com/go-ozzo/ozzo-validation/v4"

type BotRecord struct {
	Id int64 `json:"id" example:"1"`
	Bot
}

type Bot struct {
	Name   string  `json:"name"    example:"Регистрация на мероприятие"`
	Token  string  `json:"token"   example:"XXXXXXXXXX:YYYYYYYYYYYYYYYYYYYYYYYYYYYYYYYYYYY"`
	Start  int64   `json:"start"   example:"1"`
	Blocks []Block `json:"blocks"`
}

func (b Bot) Validate() error {
	return validation.ValidateStruct(&b,
		validation.Field(&b.Name, validation.Required),
		validation.Field(&b.Token, validation.Required),
		validation.Field(&b.Start, validation.Required, validation.Min(1)),
		validation.Field(&b.Blocks, validation.Required, validation.Length(1, 0)),
	)
}
