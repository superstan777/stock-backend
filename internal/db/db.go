package db

import (
	"database/sql"
	"fmt"
	"os"
	"time"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func Connect() error {
	connStr := os.Getenv("DATABASE_URL")
	if connStr == "" {
		return fmt.Errorf("DATABASE_URL is not set")
	}

	var err error
	for i := 0; i < 10; i++ { // próbujemy 10 razy
		DB, err = sql.Open("postgres", connStr)
		if err == nil {
			err = DB.Ping()
		}
		if err == nil {
			fmt.Println("Connected to database successfully")
			return nil
		}
		fmt.Println("Waiting for DB to be ready...")
		time.Sleep(2 * time.Second)
	}

	return fmt.Errorf("could not connect to DB: %v", err)
}