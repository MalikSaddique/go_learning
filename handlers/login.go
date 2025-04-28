package handler

// func HandleLogin(c *gin.Context) {
// 	var u models.UserLogin
// 	if err := c.ShouldBindJSON(&u); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
// 		return
// 	}

// 	token, refreshToken, err := service.LoginUser(u.Email, u.Password)
// 	if err != nil {
// 		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
// 		return
// 	}

// 	c.JSON(http.StatusOK, gin.H{
// 		"AccessToken":  token,
// 		"RefreshToken": refreshToken,
// 	})
// }
