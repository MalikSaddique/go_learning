package connection

import (
	"database/sql"
	"fmt"
	"os"
	"strconv"

	_ "github.com/lib/pq"
)

func DbConnection() *sql.DB {
	host := os.Getenv("DB_HOST")
	portStr := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")

	port, err := strconv.Atoi(portStr)
	if err != nil {
		fmt.Println("Error converting port to integer:", err)
		return nil
	}

	psqlconn := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname,
	)

	db, err := sql.Open("postgres", psqlconn)
	if err != nil {
		fmt.Println("Error opening DB connection:", err)
		return nil
	}

	err = db.Ping()
	if err != nil {
		fmt.Println("Failed to connect to database:", err)
		return nil
	}

	fmt.Println("Connected to PostgreSQL successfully!")
	return db
}
