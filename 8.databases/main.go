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
	Name   string `gorm:"size:255"`
	Email  string `gorm:"not null;uniqueIndex"`
	Orders []Order
}

type Order struct {
	gorm.Model
	UserID      int  `gorm:"not null"`
	Amount      uint `gorm:"not null"`
	Description *string
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

	if err := db.AutoMigrate(User{}, Order{}); err != nil {
		panic(err)
	}

	var users []User

	if err := db.Preload("Orders").Find(&users).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			fmt.Println("User not found!")
			os.Exit(1)
		default:
			panic(err)
		}
	}

	fmt.Println(users)

	//db.Migrator().DropTable(User{})
}

var orders = []Order{
	{Amount: 100, Description: boxString("USB-C Charger"), UserID: 3},
	{Amount: 200, Description: boxString("Web Cam"), UserID: 7},
	{Amount: 300, Description: boxString("USB-C Charger"), UserID: 5},
	{Amount: 10, Description: boxString("Apple Sticker"), UserID: 11},
	{Amount: 120, Description: boxString("Airpods Refurbished"), UserID: 6},
	{Amount: 50, Description: boxString("AAA Battery 1 pack"), UserID: 7},
	{Amount: 99, Description: boxString("Note 3 Book 3x"), UserID: 8},
}

func (u User) String() string {
	return fmt.Sprintf(
		"\n(ID: %d, Name: %s, Email: %s, Orders: %v)\n",
		u.ID, u.Name, u.Email, u.Orders,
	)
}
func (o Order) String() string {
	return fmt.Sprintf(
		"(ID: %d, Amount: %d, UserID: %d, Description: %v)\n",
		o.ID, o.Amount, o.UserID, o.Description,
	)
}
func boxString(s string) *string {
	return &s
}
