package dto

import (
	"obserbooks/infrastructure"
)

type BookDTO struct {
	UserId         string
	JanCode        string `json:"jancode"`
	Title          string `json:"title"`
	Author         string `json:"author"`
	PublisherName  string `json:"publishername"`
	ItemUrl        string `json:"itemurl"`
	LargeImageUrl  string `json:"larageimageurl"`
	MediumImageUrl string `json:"mediumimageurl"`
	MyBook         bool   `json:"mybook"`
	AttentionBook  bool   `json:"attentionbook"`
}

func (b *BookDTO) TableName() string {
	return "books"
}

var _ infrastructure.DTO = &BookDTO{}
