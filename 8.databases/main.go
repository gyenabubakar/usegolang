package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
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
	Email string `gorm:"not null;uniqueIndex"`
}

func main() {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second,  // Slow SQL threshold
			LogLevel:                  logger.Error, // Log level
			IgnoreRecordNotFoundError: false,        // Ignore ErrRecordNotFound error for logger
			Colorful:                  true,         // Disable color
		},
	)

	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		panic("Failed to connect to database")
	}
	defer func() {
		sqlDB, _ := db.DB()
		if err := sqlDB.Close(); err != nil {
			panic(err)
		}
	}()

	if err := db.AutoMigrate(User{}); err != nil {
		panic(err)
	}

	//var user User
	var users []User

	db.Find(&users)

	fmt.Println(users)

	//db.Migrator().DropTable(User{})
}

//var users = []User{
//	{Name: "Sena", Email: "sena@test.io"},
//	{Name: "Felix", Email: "felix@test.io"},
//	{Name: "James", Email: "james@test.io"},
//	{Name: "Emma", Email: "emma@test.io"},
//	{Name: "Biney", Email: "biney@test.io"},
//	{Name: "Ben", Email: "ben@test.io"},
//	{Name: "Percy", Email: "percy@test.io"},
//}

func (u User) String() string {
	return fmt.Sprintf("\n(ID: %d, Name: %s, Email: %s)\n", u.ID, u.Name, u.Email)
}
