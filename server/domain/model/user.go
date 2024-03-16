package model

type User struct {
	UserId   string `gorm:"primaryKey"`
	Name     string
	Email    string
	Password string
	Books    []Book `gorm:"foreignKey:UserId;references:UserId"`
}
