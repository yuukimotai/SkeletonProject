package repositories

import (
	"fmt"
	repositories "obserbooks/domain/domain-repositories"
	usermodel "obserbooks/domain/user"
	"obserbooks/infrastructure/dto"
	database "obserbooks/infrastructure/mysql"
)

type UserRepository struct {
	db *database.DB
}

var _ repositories.UserRepository = &UserRepository{}

// NewUserRepository コンストラクタ
func NewUserRepository(db *database.DB) repositories.UserRepository {
	return &UserRepository{
		db: db,
	}
}

// Create ユーザを作成します．
func (ur *UserRepository) Create(user *usermodel.User) error {
	userDTO := dto.UserDTO{
		UserId:   user.UserId,
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
	}

	err := ur.db.Create(&userDTO)

	if err != nil {
		return err
	}

	// DTOをユーザエンティティに変換します．
	return nil
}

// FindById IDを元にユーザを返却します．
func (ur *UserRepository) FindByEmail(user usermodel.User) (*usermodel.User, error) {
	userDTO := dto.UserDTO{}

	result, err := ur.db.Find(userDTO)
	if err != nil {
		println(err)
	}

	fmt.Println(err)

	// DTOをユーザエンティティに変換します．
	return result, nil
}

// FindAll 全てのユーザを返却します．
func (ur *UserRepository) FindAll() (usermodel.Users, error) {
	usersDTO := dto.UsersDTO{}

	err := ur.db.FindAll(&usersDTO)

	if err != nil {
		return nil, err
	}

	// DTO配列をユーザエンティティ配列に変換します．
	return usersDTO.ToUsers(), nil
}

// // Update ユーザを更新します．
// func (ur *UserRepository) Update(user *usermodel.User) error {
// 	// ユーザエンティティをDTOに変換します．
// 	userDTO := dtos.UserDTO{
// 		UserId:            user.Id().ToPrimitive(),
// 		Name:      user.Name().LastName(),
// 	}

// 	err := ur.db.Updates(&userDTO)

// 	if err != nil {
// 		return err
// 	}

// 	// DTOをユーザエンティティに変換します．
// 	return nil
// }

// // Delete ユーザを削除します．
// func (ur *UserRepository) Delete(id ids.UserId) error {
// 	// ユーザエンティティをDTOに変換します．
// 	userDTO := dtos.UserDTO{
// 		UserId: id.ToPrimitive(),
// 	}

// 	err := ur.db.Delete(&userDTO)

// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }
