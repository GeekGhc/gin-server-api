package util

var jwtSecret []byte

type Claims struct {
	Username string `json:"username"`
	Password string 	`json:"password"`
}
