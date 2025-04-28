package database

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/MalikSaddique/go_learning/models"
	"github.com/gin-gonic/gin"
)

type Storage interface {
	SaveResult(*models.Result) (*models.Result, error)
	FindUserByEmail(email string) (*models.UserLogin, error)
	SignUp(c *gin.Context, req *models.User) *models.User
	FetchResultsByUserID(userID string, limit int, offset int) ([]models.Result, error)
}

type StorageImpl struct {
	db *sql.DB
}

func NewStorage(db *sql.DB) Storage {
	return &StorageImpl{
		db: db,
	}
}

func (u *StorageImpl) SaveResult(result *models.Result) (*models.Result, error) {

	query := `
		INSERT INTO results 
		(words, digits, special_char, lines, spaces, sentences, punctuation, consonants, vowels, user_id)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)`

	_, err := u.db.Exec(query,
		result.Words,
		result.Digits,
		result.SpecialChar,
		result.Lines,
		result.Spaces,
		result.Sentences,
		result.Punctuation,
		result.Consonants,
		result.Vowels,
		result.Id,
	)

	if err != nil {
		log.Println("Error inserting result:", err)
		return nil, err
	}

	fmt.Println("Result saved successfully.")
	return &models.Result{
		Words:       result.Words,
		Digits:      result.Digits,
		SpecialChar: result.SpecialChar,
		Lines:       result.Lines,
		Spaces:      result.Spaces,
		Sentences:   result.Sentences,
		Punctuation: result.Punctuation,
		Consonants:  result.Consonants,
		Vowels:      result.Vowels,
		Id:          result.Id,
	}, nil
}

func (u *StorageImpl) FindUserByEmail(email string) (*models.UserLogin, error) {
	fmt.Println(59)

	var user models.UserLogin

	err := u.db.QueryRow("SELECT id,  email, password FROM users WHERE email=$1", email).Scan(&user.Id, &user.Email, &user.Password)
	if err != nil {
		fmt.Println(err, "22")
		return nil, err
	}

	return &user, nil

}

func (u *StorageImpl) SignUp(c *gin.Context, req *models.User) *models.User {
	// var req *models.User
	err := u.db.QueryRow("SELECT email FROM users WHERE email = $1", &req.Email).Scan(&req.Email)
	if err == nil {
		c.JSON(http.StatusConflict, gin.H{"error": "User already exists with this email"})
		return nil
	} else if err != sql.ErrNoRows {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error checking existing user"})
		return nil
	}

	_, err = u.db.Exec("INSERT INTO users (email, password) VALUES ($1, $2)", &req.Email, &req.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return nil
	}

	return req

}

func (u *StorageImpl) FetchResultsByUserID(userID string, limit int, offset int) ([]models.Result, error) {
	rows, err := u.db.Query(`SELECT user_id, words, digits, special_char, lines, spaces, sentences, punctuation, consonants, vowels 
	                         FROM results WHERE user_id = $1 ORDER BY user_id LIMIT $2 OFFSET $3`, userID, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("query failed: %w", err)
	}
	defer rows.Close()

	var results []models.Result
	for rows.Next() {
		var result models.Result
		err := rows.Scan(
			&result.Id,
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
			return nil, fmt.Errorf("row scan error: %w", err)
		}
		results = append(results, result)
	}

	return results, nil
}
