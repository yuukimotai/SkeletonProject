package domain

import (
	usermodel "obserbooks/domain/user"
)

type AuthenticationRepositorys interface {
	Login(*usermodel.User) error
	Logout(*usermodel.User) (*usermodel.User, error)
}
