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

	var users []User

	rows, err := db.Query(`SELECT id, name, email FROM users`)
	if err != nil {
		panic(err)
	}

	for rows.Next() {
		var user User
		if err = rows.Scan(&user.ID, &user.Name, &user.Email); err != nil {
			panic(err)
		}
		users = append(users, user)
	}

	if rowErr := rows.Err(); rowErr != nil {
		panic(rowErr)
	}

	fmt.Println(users)
}

type User struct {
	ID    int
	Name  string
	Email string
}

func (u User) String() string {
	return fmt.Sprintf(
		"(id:%d, name:%s, email:%s)\n",
		u.ID, u.Name, u.Email,
	)
}

func closeDB(db *sql.DB) {
	if err := db.Close(); err != nil {
		panic(err)
	}
}
