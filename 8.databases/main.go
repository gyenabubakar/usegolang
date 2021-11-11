package main

import (
	"fmt"
	//"database/sql"
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
	fmt.Println(psqlInfo)
}
