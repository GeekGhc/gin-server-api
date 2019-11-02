package models

import "github.com/jinzhu/gorm"

var db *gorm.DB

// CloseDB closes database connection (unnecessary)
func CloseDB() {
	defer db.Close()
}
