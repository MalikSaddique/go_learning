package authserviceimpl

import (
	"github.com/MalikSaddique/go_learning/database"
	authservice "github.com/MalikSaddique/go_learning/service/auth_service"
)

type AuthServiceImpl struct {
	userAuth database.Storage
}

func NewUserSErviceImpl(input NewUserServiceImpl) authservice.AuthService {
	return &AuthServiceImpl{
		userAuth: input.UserAuth,
	}
}

type NewUserServiceImpl struct {
	UserAuth database.Storage
}
