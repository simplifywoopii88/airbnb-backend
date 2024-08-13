package database

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
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

	// CreateEnumType
	CreateCityType(db)
	CreateKindType(db)

	log.Println("Connected to the database successfully")
	log.Println("Running Migrations")

	if err := db.AutoMigrate(
		new(User),
		new(Room),
	); err != nil {
		log.Println("cannot migrate DB")
		panic(err.Error())
	}

	DB = db
}
