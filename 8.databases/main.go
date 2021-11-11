package main

import (
	"database/sql"
	"fmt"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"os"
)

var dbURL string

func init() {
	if err := godotenv.Load(); err != nil {
		panic(err)
		return
	}
	dbURL = os.Getenv("DB_URL")
}

func main() {
	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		panic(err)
	}
	defer closeDB(db)

	var id int
	err = db.QueryRow(`
		INSERT INTO users(name, email)
		VALUES($1,$2)
		RETURNING id`,
		"Felix", "felix@office.org",
	).Scan(&id)
	if err != nil {
		panic(err)
	}

	fmt.Println("ID is...", id)
}

//type User struct {
//
//}

func closeDB(db *sql.DB) {
	if err := db.Close(); err != nil {
		panic(err)
	}
}
