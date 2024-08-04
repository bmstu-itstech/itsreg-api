package model

import "github.com/go-ozzo/ozzo-validation/v4"

type BlockType string

const (
	Message   BlockType = "message"
	Question  BlockType = "question"
	Selection BlockType = "selection"
)

type Block struct {
	State   int64     `json:"state"    example:"1"`
	Type    BlockType `json:"type"     example:"question"`
	Default int64     `json:"default"  example:"2"`
	Title   string    `json:"title"    example:"ФИО участника"`
	Text    string    `json:"text"     example:"Введите своё ФИО, например, Иванов Иван Иванович"`
	Options []Option  `json:"options"`
}

func (b Block) Validate() error {
	return validation.ValidateStruct(&b,
		validation.Field(&b.State, validation.Required, validation.Min(1)),
		validation.Field(&b.Type, validation.Required, validation.In(Message, Question, Selection)),
		validation.Field(&b.Default, validation.Required, validation.Min(0)),
		validation.Field(&b.Title, validation.Required),
		validation.Field(&b.Text, validation.Required),
	)
}
