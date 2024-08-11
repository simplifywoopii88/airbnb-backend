package database

import (
	"fmt"
	"github.com/simplifywoopii88/airbnb-backend/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB

const (
	host     = "localhost"
	user     = "postgres"
	password = "test123!@#"
	dbname   = "postgres"
	port     = "5432"
	sslmode  = "disable"
	TimeZone = "Asia/Seoul"
)

func ConnectDB() {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
		host,
		user,
		password,
		dbname,
		port,
		sslmode,
		TimeZone,
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Databse base error, %v\n", err)
	}

	log.Println("Connected to the database successfully")
	log.Println("Running Migrations")

	if err := db.AutoMigrate(new(models.User)); err != nil {
		log.Println("cannot migrate DB")
		panic(err.Error())
	}

	DB = db
}
