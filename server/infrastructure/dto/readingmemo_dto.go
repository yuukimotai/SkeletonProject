package dto

import (
	"obserbooks/infrastructure"
)

type ReadingMemoDTO struct {
	MemoId  string `json:"memoid"`
	UserId  string
	JanCode string `json:"jancode"`
	Title   string `json:"title"`
	Author  string `json:"author"`
	Content string `json:"content"`
	//Status     string `json:"status"`
}

func (r *ReadingMemoDTO) TableName() string {
	return "reading_memos"
}

var _ infrastructure.DTO = &ReadingMemoDTO{}
