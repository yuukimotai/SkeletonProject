package database

import (
	"errors"
	"fmt"
	"obserbooks/domain/model"
	usermodel "obserbooks/domain/user"
	"obserbooks/infrastructure"
	"obserbooks/infrastructure/dto"

	"gorm.io/gorm"
)

func (d *DB) Create(DTO infrastructure.DTO) error {
	return d.connection.Create(DTO).Error
}

func (d *DB) FindAll(DTO infrastructure.DTO) error {
	return d.connection.First(DTO).Error
}

func (d *DB) Find(user dto.UserDTO) (*usermodel.User, error) {
	var target *usermodel.User
	err := d.connection.Where("email = ?", user.Email).Or("password = ?", user.Password).Find(&target).Error
	fmt.Println(user.UserId)
	if err != nil {
		fmt.Println(err)
	}
	return target, err
}

func (d *DB) FindByEmailAndPassword(user dto.UserDTO) *model.User {
	target := &model.User{}
	err := d.connection.Where("email = ?", user.Email).Or("password = ?", user.Password).Find(&target).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		fmt.Println("該当レコードなし")
	}
	return target
}
