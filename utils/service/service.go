package service

import (
	"github.com/jinzhu/gorm"
	"strconv"
)

func Limit(limit string) int {
	limitInt, err := strconv.Atoi(limit)
	// if user not inputted the value,
	// set limit at 25
	if err != nil {
		limitInt = 25
	}
	return limitInt
}

func Search(search string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if search != "" {
			db = db.Where("name LIKE ?", "%"+search+"%")
		}
		return db
	}
}

func Offset(offset string) int {
	offsetInt, err := strconv.Atoi(offset)
	if err != nil {
		offsetInt = 0
	}
	return offsetInt
}

// function for skipping an offset
// WIP: create a method in interfaces
func Skip(skip int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if skip > 0 {
			db = db.Offset(skip)
		}
		return db
	}
}