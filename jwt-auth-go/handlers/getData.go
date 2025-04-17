package handler

import (
	"fmt"

	"github.com/MalikSaddique/go_learning/jwt-auth-go/analyzer"
	connection "github.com/MalikSaddique/go_learning/jwt-auth-go/database"
	"github.com/gin-gonic/gin"
)

func GetResults(c *gin.Context) {
	id := c.Param("user_id")

	db, err := connection.DbConnection()
	if err != nil {
		fmt.Println("Database connection error:", err)
		c.JSON(500, gin.H{"error": "Database connection failed"})
		return
	}
	defer db.Close()
	// var user UserInfo
	var result analyzer.Result

	row := db.QueryRow("SELECT words, digits, special_char, lines, spaces, sentences, punctuation, consonants, vowels FROM results WHERE user_id = $1", id)

	err = row.Scan(&result.Words, &result.Digits, &result.SpecialChar, &result.Lines, &result.Spaces, &result.Sentences, &result.Punctuation, &result.Consonants, &result.Vowels)
	if err != nil {
		fmt.Println("Error retrieving results:", err)
		c.JSON(404, gin.H{"error": "Result not found"})
		return
	}

	c.JSON(200, gin.H{
		"words":        result.Words,
		"digits":       result.Digits,
		"special_char": result.SpecialChar,
		"lines":        result.Lines,
		"spaces":       result.Spaces,
		"sentences":    result.Sentences,
		"punctuation":  result.Punctuation,
		"consonants":   result.Consonants,
		"vowels":       result.Vowels,
	})
}
