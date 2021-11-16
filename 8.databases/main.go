package main

import (
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
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

type User struct {
	gorm.Model
	Name  string `gorm:"size:255"`
	Email string `gorm:"unique;uniqueIndex"`
}

func main() {
	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})
	if err != nil {
		panic("failed to connect to database")
	}
	sqlDB, _ := db.DB()
	if err := sqlDB.Ping(); err != nil {
		panic(err)
	}

	if err := db.AutoMigrate(User{}); err != nil {
		panic(err)
	}
}
