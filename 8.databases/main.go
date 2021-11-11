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

	if err := db.Ping(); err != nil {
		panic(err)
	}

	fmt.Println("Connection to DB established!")

	if err := db.Close(); err != nil {
		panic(err)
	}
}
