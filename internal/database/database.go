package database

import (
	"fmt"
	"time"

	"github.com/ianprogrammer/go-api-integration-test/internal/configuration"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDatabase(database configuration.DatabaseConfig) (*gorm.DB, error) {

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d", database.Host, database.UserName, database.Password, database.DatabaseName, database.DatabasePort)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	sqlDB, err := db.DB()

	if err != nil {
		return nil, err
	}

	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

	return db, nil

}
