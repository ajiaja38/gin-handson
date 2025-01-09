package config

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func DBConnection() (*gorm.DB, error) {
	host := os.Getenv("DB_HOST")
	pass := os.Getenv("DB_PASSWORD")
	user := os.Getenv("DB_USER")
	port := os.Getenv("DB_PORT")
	name := os.Getenv("DB_NAME")

	dbURI := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, pass, name)
	db, err := gorm.Open(postgres.Open(dbURI), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		return nil, err
	}

	return db, nil
}
