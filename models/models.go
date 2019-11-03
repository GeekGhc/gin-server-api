package models

import "github.com/jinzhu/gorm"

var db *gorm.DB

type Model struct {
	ID        int `gorm:"primary_key" json:"id"`
	CreatedAt int `json:"created_at"`
	UpdatedAt int `json:"updated_at"`
	DeleteAt  int `json:"deleted_at"`
}

// initializes the database instance
func Setup(){

}


// CloseDB closes database connection (unnecessary)
func CloseDB() {
	defer db.Close()
}
