package dto

import (
	usermodel "obserbooks/domain/user"
	"obserbooks/infrastructure"

	"github.com/google/uuid"
)

// UserDTO NOTE: 利便性のため，DTOはパブリックフィールドとします．
type UserDTO struct {
	UserId   uuid.UUID `gorm:"primaryKey"`
	Name     string
	Email    string
	Password string
}

type UsersDTO []*UserDTO

var _ infrastructure.DTO = &UserDTO{}

// TableName テーブル名を定義します．
func (ud UserDTO) TableName() string {
	return "users"
}

// ToUser DTOをユーザエンティティに変換します．
func (ud UserDTO) ToUser() *usermodel.User {
	user := usermodel.NewUser(usermodel.User{
		UserId:   ud.UserId,
		Name:     ud.Name,
		Email:    ud.Email,
		Password: ud.Password,
	})
	return user
}

// ToUsers UsersDTOをUsersに変換します．
func (usd UsersDTO) ToUsers() usermodel.Users {
	users := usermodel.Users{}

	for i, ud := range usd {
		users[i] = ud.ToUser()
	}

	return users
}
