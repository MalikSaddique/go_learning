package controllers

import (
	"github.com/gin-gonic/gin"
)

func Data(c *gin.Context) {
	c.JSON(200, gin.H{
		"name":     "Siddique",
		"Position": "Assistant",
		"Email":    "maliksaddique139@gmail.com",
		"password": "abcd123",
	})
}
