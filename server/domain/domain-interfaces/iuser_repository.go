package domain

import (
	usermodel "obserbooks/domain/user"
)

type UserRepository interface {
	Create(*usermodel.User) error
	FindByEmail(*usermodel.User) (*usermodel.User, error)
}
