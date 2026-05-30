package link

import (
	"kilkenny/purpleschool/pkg/helpers"

	"gorm.io/gorm"
)

type Link struct {
	gorm.Model
	Url  string `json:"url" `
	Hash string `json:"hash" gorm:"uniqueIndex"`
}

func NewLink(url string) *Link {
	return &Link{
		Url:  url,
		Hash: helpers.GenerateStringRune(16),
	}
}
