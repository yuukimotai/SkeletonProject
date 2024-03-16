package model

type ReadingMemo struct {
	MemoId  string
	UserId  string
	JanCode string
	Title   string
	Author  string
	Content string
}

type NewReadingMemo struct {
	MemoId  string
	UserId  string
	JanCode string
	Title   string
	Author  string
	Content string
}

type GetMemo struct {
	MemoId  string
	UserId  string
	JanCode string
}

type GetAllMemo struct {
	UserId  string
	JanCode string
}

type UpdateReadingMemo struct {
	MemoId  string
	UserId  string
	JanCode string
	Title   string
	Author  string
	Content string
}

type DeleteReadingMemo struct {
	MemoId  string
	UserId  string
	JanCode string
}
