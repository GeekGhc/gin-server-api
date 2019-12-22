package models

import "github.com/jinzhu/gorm"

type User struct {
	Model
	Username string `json:"username"`
	Email    string `gorm:"type:varchar(255)" json:"email"`
	Password string `json:"password"`
}

// check auth if authentication information exist
func CheckAuth(username, password string) (bool, error) {
	var user User
	err := db.Select("id").Where(User{Username: username, Password: password}).First(&user).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}

	if user.ID > 0 {
		return true, nil
	}

	return false, nil
}
