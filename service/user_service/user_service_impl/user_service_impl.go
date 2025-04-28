package userserviceimpl

import (
	"github.com/MalikSaddique/go_learning/database"
	userservice "github.com/MalikSaddique/go_learning/service/user_service"
)

type UserServiceImpl struct {
	UserAuth database.Storage
}

func NewUserService(input database.Storage) userservice.UserService {
	return &UserServiceImpl{
		UserAuth: input,
	}
}

type NewUserServiceImpl struct {
	UserAuth database.Storage
}
