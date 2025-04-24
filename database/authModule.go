package database

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/MalikSaddique/go_learning/analyzer"
	"github.com/MalikSaddique/go_learning/models"
)

type Storage interface {
	SaveResult(analyzer.Result) error
	FindUserByEmail(email string) (*models.UserLogin, error)
	// LoginUser(email, password string) (string, string, error)
}

type StorageImpl struct {
	db *sql.DB
}

func NewStorage(db *sql.DB) Storage {
	return &StorageImpl{
		db: db,
	}
}

// var _ Storage = &StorageImpl{}
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

	var user models.UserLogin

	err := u.db.QueryRow("SELECT id,  email, password FROM users WHERE email=$1", email).Scan(&user.Id, &user.Email, &user.Password)
	if err != nil {
		return nil, err
	}

	return &user, nil

}
