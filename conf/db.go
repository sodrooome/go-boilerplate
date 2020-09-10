package conf

import (
	"backend-project/domain"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

// var db *gorm.DB
var product domain.Products

func ValidateDB() (*gorm.DB, error) {
	db, err := gorm.Open("sqlite3", "./products.db")
	if err != nil {
		panic("Can't connect the database")
	}

	db.Model(product).Rows()
	db.AutoMigrate(product)

	return db, nil
}