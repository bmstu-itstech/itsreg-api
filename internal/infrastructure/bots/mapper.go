package bots

import (
	"github.com/bmstu-itstech/itsreg-api/internal/model"
	botsv1 "github.com/bmstu-itstech/itsreg-proto/gen/go/bots"
)

func optionToPb(option model.Option) *botsv1.BlockOption {
	return &botsv1.BlockOption{
		Text: option.Text,
		Next: option.Next,
	}
}

func optionsToPb(options []model.Option) []*botsv1.BlockOption {
	res := make([]*botsv1.BlockOption, len(options))
	for i, option := range options {
		res[i] = optionToPb(option)
	}
	return res
}

func blockTypeToPb(s model.BlockType) botsv1.BlockType {
	switch s {
	case model.Message:
		return botsv1.BlockType_BlockMessage
	case model.Question:
		return botsv1.BlockType_BlockQuestion
	case model.Selection:
		return botsv1.BlockType_BlockSelection
	default:
		return botsv1.BlockType_Unknown
	}
}

func blockToPb(block model.Block) *botsv1.Block {
	return &botsv1.Block{
		State:   block.State,
		Type:    blockTypeToPb(block.Type),
		Default: block.Default,
		Title:   block.Title,
		Text:    block.Text,
		Options: optionsToPb(block.Options),
	}
}

func blocksToPb(blocks []model.Block) []*botsv1.Block {
	res := make([]*botsv1.Block, len(blocks))
	for i, block := range blocks {
		res[i] = blockToPb(block)
	}
	return res
}
