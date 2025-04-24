package authserviceimpl

import (
	"github.com/MalikSaddique/go_learning/models"
	"github.com/gin-gonic/gin"
)

// type TokenPair struct {
// 	AccessToken  string `json:"access_token"`
// 	RefreshToken string `json:"refresh_token"`
// }

func (u *AuthServiceImpl) Login(c *gin.Context, req *models.UserLogin) (*models.TokenPair, error) {
	token, refreshToken, err := u.db.LoginUser(req.Email, req.Password)
	if err != nil {
		return nil, err
	}

	response := models.TokenPair{
		AccessToken:  token,
		RefreshToken: refreshToken,
	}
	return &response, nil
}
