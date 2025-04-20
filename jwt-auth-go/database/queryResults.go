package connection

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/MalikSaddique/go_learning/jwt-auth-go/analyzer"
)

func SaveResult(db *sql.DB, result analyzer.Result) error {

	query := `
		INSERT INTO results 
		(words, digits, special_char, lines, spaces, sentences, punctuation, consonants, vowels, user_id)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)`

	_, err := db.Exec(query,
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
