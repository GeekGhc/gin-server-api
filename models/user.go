package models

import (
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Model
	Username string `json:"username"`
	Email    string `gorm:"type:varchar(255)" json:"email"`
	Password string `json:"password"`
}

// check auth if authentication information exist
func CheckAuth(username, password string) (bool, error) {
	var user User
	err := db.Where("username = ?", username).First(&user).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}
	//密码校验
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return false, err
	}

	if user.ID > 0 {
		return true, nil
	}

	return false, nil
}
