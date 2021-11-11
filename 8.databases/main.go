package main

import (
	"database/sql"
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

	_, err = db.Exec(
		`INSERT INTO users(name, email) VALUES($1,$2)`,
		"Gyen", "gyen@dev.co",
	)
	if err != nil {
		panic(err)
	}
}

func closeDB(db *sql.DB) {
	if err := db.Close(); err != nil {
		panic(err)
	}
}
