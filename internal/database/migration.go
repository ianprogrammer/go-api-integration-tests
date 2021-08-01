package database

import (
	"github.com/ianprogrammer/go-api-integration-test/pkg/product"
	"gorm.io/gorm"
)

func MigrateDB(db *gorm.DB) error {

	if err := db.AutoMigrate(&product.Product{}); err != nil {
		return err
	}
	return nil
}
