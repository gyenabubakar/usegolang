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
	var name string
	var email string

	err = db.QueryRow(`
		SELECT id, name, email FROM users WHERE id=$1`,
		5,
	).Scan(&id, &name, &email)
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("No records found.")
		} else {
			panic(err)
		}
	}

	fmt.Printf(
		"id: %d\t|\tname: %s\t|\temail:%s\n",
		id, name, email,
	)
}

//type User struct {
//
//}

func closeDB(db *sql.DB) {
	if err := db.Close(); err != nil {
		panic(err)
	}
}
