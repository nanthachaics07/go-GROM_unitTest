package main

import (
	"errors"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "myuser"
	password = "mypassword"
	dbname   = "mydatabase"
)

type User struct {
	gorm.Model
	Fullname string
	Email    string `gorm:"unique"`
	Age      int
}

func InitializeDB() *gorm.DB {
	// Configure your PostgreSQL database details here
	dsn := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect to database")
	}
	db.AutoMigrate(&User{})
	return db
}

func AddUser(db *gorm.DB, fullname, email string, age int) error {
	user := User{Fullname: fullname, Email: email, Age: age}

	var count int64
	db.Model(&User{}).Where("email = ?", email).Count(&count)
	if count > 0 {
		return errors.New("email already exists")
	}

	result := db.Create(&user)
	return result.Error
}

func main() {
	db := InitializeDB()

	AddUser(db, "Jhon Zenaaa", "jane.smitttt@test.com", 22)
}
