package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

var DB *gorm.DB

func ConnectDataBase() {

	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	Dbdriver := os.Getenv("DB_DRIVER")
	DbHost := os.Getenv("DB_HOST")
	DbPort := os.Getenv("DB_PORT")
	DbUser := os.Getenv("DB_USER")
	DbPassword := os.Getenv("DB_PASSWORD")
	DbName := os.Getenv("DB_NAME")

	DBURL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", DbUser, DbPassword, DbHost, DbPort, DbName)

	DB, err = gorm.Open(Dbdriver, DBURL)

	if err != nil {
		fmt.Println("Cannot connect to database", Dbdriver)
		log.Fatal("connection error:", err)
	} else {
		fmt.Println("Connected to database", Dbdriver)
	}

	DB.AutoMigrate(&User{})
}