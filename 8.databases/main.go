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

	var orders []Order

	rows, err := db.Query(`
		SELECT users.id,
				users.name,
				users.email,
				orders.id,
				orders.amount,
				orders.description
		FROM users
				INNER JOIN orders on users.id = orders.user_id
	`)
	if err != nil {
		panic(err)
	}
	for rows.Next() {
		var order Order
		err := rows.Scan(
			&order.UserID, &order.UserName, &order.UserEmail, &order.ID, &order.Amount, &order.Description,
		)
		if err != nil {
			panic(err)
		}
		orders = append(orders, order)
	}
	if rows.Err() != nil {
		panic(rows.Err())
	}

	fmt.Printf("%#v\n", orders)
}

type Order struct {
	ID          int
	Amount      int
	Description string
	UserID      int
	UserName    string
	UserEmail   string
}

func closeDB(db *sql.DB) {
	if err := db.Close(); err != nil {
		panic(err)
	}
}
