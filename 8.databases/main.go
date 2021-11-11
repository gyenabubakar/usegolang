package main

import (
	"database/sql"
	"fmt"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"os"
)

var (
	dbHost     string
	dbHostPort string
	dbUser     string
	dbPassword string
	dbName     string
)

func init() {
	if err := godotenv.Load(); err != nil {
		panic(err)
		return
	}
	dbHost = os.Getenv("DB_HOST")
	dbHostPort = os.Getenv("DB_HOST_PORT")
	dbUser = os.Getenv("DB_USER")
	dbPassword = os.Getenv("DB_PASSWORD")
	dbName = os.Getenv("DB_NAME")
}

func main() {
	psqlInfo := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disabled",
		dbHost, dbHostPort, dbUser, dbPassword, dbName,
	)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	fmt.Println("Connection to DB established!")

	if err := db.Close(); err != nil {
		panic(err)
	}
}
