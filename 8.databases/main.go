package main

import (
	"bufio"
	"fmt"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"strings"
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
		panic("failed to connect to database")
	}
	sqlDB, _ := db.DB()
	if err := sqlDB.Ping(); err != nil {
		panic(err)
	}
	if err := db.AutoMigrate(User{}); err != nil {
		panic(err)
	}

	name, email := getInfo()
	user := User{
		Name:  name,
		Email: email,
	}
	if err := db.Create(&user).Error; err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("%+v\n", user)

	//db.Migrator().DropTable(User{})
}

func getInfo() (name, email string) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Your name?")
	name, _ = reader.ReadString('\n')
	name = strings.TrimSpace(name)

	fmt.Println("Your email?")
	email, _ = reader.ReadString('\n')
	email = strings.TrimSpace(email)
	return
}
