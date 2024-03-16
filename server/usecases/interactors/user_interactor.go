package interactors

import (
	"crypto/sha256"
	"encoding/hex"
	repositories "obserbooks/domain/domain-repositories"
	usermodel "obserbooks/domain/user"
	"obserbooks/usecases/boundaries"
	"obserbooks/usecases/requests"
	"obserbooks/usecases/responses"

	"github.com/google/uuid"
)

// TODO
// Casbin使って管理権限者しか使えないようにする　それができるまではルーター公開しない
type UserInteractor struct {
	userRepository repositories.UserRepository
}

var _ boundaries.UserInputBoundary = &UserInteractor{}

// NewUserInteractor コンストラクタ
func NewUserInteractor(userRepository repositories.UserRepository) *UserInteractor {
	return &UserInteractor{
		userRepository: userRepository,
	}
}

// CreateUser ユーザを作成します．
func (ui *UserInteractor) CreateUser(userCreateRequest *requests.UserCreateRequest) (*responses.UserCreateResponse, error) {
	var request usermodel.User

	request.UserId, _ = uuid.NewUUID()
	request.Name = userCreateRequest.Name
	request.Email = userCreateRequest.Email

	p := []byte(userCreateRequest.Password)
	sha256 := sha256.Sum256(p)
	request.Password = hex.EncodeToString(sha256[:])

	user := usermodel.NewUser(request)

	err := ui.userRepository.Create(user)

	if err != nil {
		return nil, err
	}

	return &responses.UserCreateResponse{
		UserId: user.UserId,
		Name:   user.Name,
		Email:  user.Email,
	}, nil
}

// GetUser ユーザを取得します．
func (ui *UserInteractor) GetUser(userGetRequest *requests.UserGetRequest) (*responses.UserGetResponse, error) {
	var request usermodel.User

	request.Name = userGetRequest.Name
	request.Email = userGetRequest.Email

	user, err := ui.userRepository.FindByEmail(request)

	if err != nil {
		return nil, err
	}

	return &responses.UserGetResponse{
		Name:  user.Name,
		Email: user.Email,
	}, nil
}
