package authserviceimpl

import (
	"github.com/MalikSaddique/go_learning/database"
	authservice "github.com/MalikSaddique/go_learning/service/auth_service"
)

// type AuthService interface {
// }

type AuthServiceImpl struct {
	db       database.Storage
	userAuth database.Storage
}

func NewUserSErviceImpl(input NewUserServiceImpl) authservice.AuthService {
	return &AuthServiceImpl{
		db:       input.Db,
		userAuth: input.UserAuth,
	}
}

type NewUserServiceImpl struct {
	Db       database.Storage
	UserAuth database.Storage
}
