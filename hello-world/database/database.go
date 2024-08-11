package database

import (
	"fmt"
	"github.com/simplifywoopii88/airbnb-backend/models"
	"github.com/simplifywoopii88/airbnb-backend/utils"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
)

type DbInstance struct {
	DB *gorm.DB
}

var Database DbInstance

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
	db.Logger = logger.Default.LogMode(logger.Info)
	log.Println("Running Migrations")
	// TODO: Add migrations
	if err := db.AutoMigrate(new(models.User), new(models.Product), new(models.Order)); err != nil {
		utils.HandleErr(err)
	}

	Database = DbInstance{
		DB: db,
	}
}
