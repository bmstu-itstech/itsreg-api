package response

import "github.com/bmstu-itstech/itsreg-api/internal/model"

type Create struct {
	BotId int64 `json:"bot_id" example:"1"`
}

type Bot struct {
	model.BotRecord
}
