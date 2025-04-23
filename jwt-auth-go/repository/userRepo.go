package repository

import (
	"github.com/MalikSaddique/go_learning/jwt-auth-go/database"
	"github.com/MalikSaddique/go_learning/jwt-auth-go/models"
)

func FindUserByEmail(email string) (*models.UserLogin, error) {
	db, err := database.DbConnection()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	var user models.UserLogin

	err = db.QueryRow("SELECT id,  email, password FROM users WHERE email=$1", email).Scan(&user.Id, &user.Email, &user.Password)
	if err != nil {
		return nil, err
	}

	return &user, nil

}
