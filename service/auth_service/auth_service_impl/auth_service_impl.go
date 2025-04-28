package authserviceimpl

import (
	"github.com/MalikSaddique/go_learning/database"
	authservice "github.com/MalikSaddique/go_learning/service/auth_service"
)

type AuthServiceImpl struct {
	userAuth database.Storage
}

func NewAuthService(input NewAuthServiceImpl) authservice.AuthService {
	return &AuthServiceImpl{
		userAuth: input.UserAuth,
	}
}

type NewAuthServiceImpl struct {
	UserAuth database.Storage
}
