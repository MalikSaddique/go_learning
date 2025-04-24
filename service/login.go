package service

import (
	"errors"
	"time"

	"github.com/MalikSaddique/go_learning/auth"
	"github.com/MalikSaddique/go_learning/database"
	"github.com/MalikSaddique/go_learning/models"

	// "github.com/MalikSaddique/go_learning/repository"
	"github.com/golang-jwt/jwt"
)

var user models.User
var refreshSecretKey = []byte("my_refresh_secret_key")

func LoginUser(email, password string) (string, string, error) {
	user, err := database.Storage.FindUserByEmail(email)
	if err != nil {
		return "", "", errors.New("User not found")
	}

	if user.Password != password {
		return "", "", errors.New("Invalid credentials")
	}

	token, err := auth.CreateToken(user.Email, int(user.Id))
	if err != nil {
		return "", "", errors.New("Failed to generate access token")
	}

	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": user.Email,
		"exp":   time.Now().Add(7 * 24 * time.Hour).Unix(),
	})

	refreshTokenString, err := refreshToken.SignedString(refreshSecretKey)
	if err != nil {
		return "", "", errors.New("Failed to generate refresh token")
	}

	return token, refreshTokenString, nil
}
