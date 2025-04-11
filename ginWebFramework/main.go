package main

import (
	"io"

	"github.com/MalikSaddique/go_learning/ginWebFramework/analyzer"
	"github.com/MalikSaddique/go_learning/ginWebFramework/controllers"
	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()

	// Define a GET route
	r.GET("/hello", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello, World!",
			"name":    "Muhammad siddique",
		})
	})

	// using parameters

	r.GET("/hy/:id/:name/:position", func(c *gin.Context) {
		var id = c.Param("id")
		var name = c.Param("name")
		var position = c.Param("position")

		c.JSON(200, gin.H{
			"user_id":       id,
			"user_name":     name,
			"user_position": position,
		})
	})

	//Post method using struct

	r.POST("/post", func(c *gin.Context) {
		type myRequest struct {
			Email    string `json:"email" binding:"required"`
			Password string `json:"password" binding:"required"`
			Gender   string `json:"gender" binding:"required"`
		}

		var request myRequest

		if err := c.BindJSON(&request); err != nil {
			c.JSON(400, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.JSON(200, gin.H{
			"email ":    request.Email,
			"password ": request.Password,
			"gender":    request.Gender,
		})

	})

	//Put method

	r.PUT("/post", func(c *gin.Context) {
		type myRequest struct {
			Email    string `json:"email" binding:"required"`
			Password string `json:"password" binding:"required"`
			Gender   string `json:"gender" binding:"required"`
		}

		var request myRequest

		if err := c.BindJSON(&request); err != nil {
			c.JSON(400, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.JSON(200, gin.H{
			"email ":    request.Email,
			"password ": request.Password,
			"gender":    request.Gender,
		})

	})

	//Patch method

	r.PATCH("/post", func(c *gin.Context) {
		type myRequest struct {
			Email    string `json:"email" binding:"required"`
			Password string `json:"password" binding:"required"`
			Gender   string `json:"gender" binding:"required"`
		}

		var request myRequest

		if err := c.BindJSON(&request); err != nil {
			c.JSON(400, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.JSON(200, gin.H{
			"email ":    request.Email,
			"password ": request.Password,
			"gender":    request.Gender,
		})

	})

	//Delete method

	r.DELETE("/post/:id", func(c *gin.Context) {
		var id = c.Param("id")

		c.JSON(200, gin.H{
			"id":      id,
			"message": "Data deleted",
		})
	})

	//Get method using controllers

	r.GET("/welcome", controllers.Hello)

	r.GET("/data", controllers.Data)

	//posting the text file

	r.POST("/upload", func(c *gin.Context) {
		file, err := c.FormFile("file")
		if err != nil {
			c.JSON(400, gin.H{
				"error": "Error occur",
			})
		}
		src, err := file.Open()
		if err != nil {
			c.JSON(400, gin.H{
				"error": "File not opened",
			})
		}
		defer src.Close()
		if err != nil {
			c.JSON(400, gin.H{
				"error": "File not closed",
			})
		}

		content, err := io.ReadAll(src)
		if err != nil {
			c.JSON(400, gin.H{
				"error": "File not read",
			})
		}

		c.JSON(200, gin.H{
			"file_name": file.Filename,
			"content":   string(content),
		})

	})

	//posting Task

	r.GET("/analyze", func(c *gin.Context) {
		result, err := analyzer.AnalyzeFile("Dummy_text.txt")
		if err != nil {
			c.JSON(500, gin.H{"error": "Failed to read file"})
			return
		}
		c.JSON(200, result)
	})

	r.Run(":8080")
}
