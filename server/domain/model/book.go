package model //graphqlのバインドも兼ねてます

type Book struct {
	ID             int    `gorm:"primaryKey;autoIncrement"`
	UserId         string `gorm:"primaryKey"`
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

type NewBook struct {
	UserId         string `gorm:"primaryKey"`
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

type FindBook struct {
	UserId  string
	JanCode string
}

type UpdateBook struct {
	UserId         string `gorm:"primaryKey"`
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

type DeleteBook struct {
	UserId  string
	JanCode string
}
