package database

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/MalikSaddique/go_learning/analyzer"
	"github.com/MalikSaddique/go_learning/models"
	"github.com/gin-gonic/gin"
)

type Storage interface {
	SaveResult(analyzer.Result) error
	FindUserByEmail(email string) (*models.UserLogin, error)
	SignUp(c *gin.Context, req *models.User) *models.User
}

type StorageImpl struct {
	db *sql.DB
}

func NewStorage(db *sql.DB) Storage {
	return &StorageImpl{
		db: db,
	}
}

func (u *StorageImpl) SaveResult(result analyzer.Result) error {

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
		result.UserID,
	)

	if err != nil {
		log.Println("Error inserting result:", err)
		return err
	}

	fmt.Println("Result saved successfully.")
	return nil
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
