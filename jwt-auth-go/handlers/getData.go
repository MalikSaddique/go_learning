package handler

import (
	"fmt"
	"strconv"

	connection "github.com/MalikSaddique/go_learning/jwt-auth-go/database"
	"github.com/MalikSaddique/go_learning/jwt-auth-go/models"
	"github.com/gin-gonic/gin"
)

// GetResults godoc
// @Summary      Get Analysis Result (Paginated)
// @Description  Retrieves paginated analyzed text data from the database for the given user ID
// @Tags         protected
// @Security     BearerAuth
// @Produce      json
// @Param        user_id path string true "User ID"
// @Param        page query int false "Page number"
// @Param        limit query int false "Limit per page"
// @Success      200  "Result data retrieved"
// @Failure      404  "Result not found"
// @Failure      500  "Database connection failed"
// @Router       /getdata/{user_id} [get]
func GetResults(c *gin.Context) {
	id := c.Param("user_id")

	pageStr := c.DefaultQuery("page", "1")
	limitStr := c.DefaultQuery("limit", "10")

	page, err := strconv.Atoi(pageStr)
	if err != nil || page < 1 {
		page = 1
	}
	limit, err := strconv.Atoi(limitStr)
	if err != nil || limit < 1 {
		limit = 10
	}
	offset := (page - 1) * limit

	db, err := connection.DbConnection()
	if err != nil {
		fmt.Println("Database connection error:", err)
		c.JSON(500, gin.H{"error": "Database connection failed"})
		return
	}
	defer db.Close()

	rows, err := db.Query(`SELECT user_id, words, digits, special_char, lines, spaces, sentences, punctuation, consonants, vowels 
	                       FROM results WHERE user_id = $1 ORDER BY user_id LIMIT $2 OFFSET $3`, id, limit, offset)
	if err != nil {
		fmt.Println("Error querying results:", err)
		c.JSON(500, gin.H{"error": "Query failed"})
		return
	}
	defer rows.Close()

	// id = models.User
	var results []models.Result
	for rows.Next() {
		var result models.Result
		err := rows.Scan(
			&result.ID,
			&result.Words,
			&result.Digits,
			&result.SpecialChar,
			&result.Lines,
			&result.Spaces,
			&result.Sentences,
			&result.Punctuation,
			&result.Consonants,
			&result.Vowels,
		)
		if err != nil {
			fmt.Println("Row scan error:", err)
			continue
		}
		results = append(results, result)
	}

	if len(results) == 0 {
		c.JSON(404, gin.H{"error": "No results found"})
		return
	}

	c.JSON(200, gin.H{
		"page":    page,
		"limit":   limit,
		"results": results,
	})
}
