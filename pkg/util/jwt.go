package util

import (
	"gin-server-api/models"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var jwtSecret []byte

type Claims struct {
	Username string `json:"username"`
	Password string `json:"password"`
	jwt.StandardClaims
}

type AuthToken struct {
	Token     string `json:"token"`
	ExpiresAt string `json:"expires_at"`
}

type userStdClaims struct {
	jwt.StandardClaims
	*models.User
}

// GenerateToken generate tokens used for auth
func GenerateToken(m *models.User) (*AuthToken, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(24 * time.Hour)

	claims := userStdClaims{
		User: m,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			IssuedAt:  time.Now().Unix(),
			Issuer:    "gin-server-api",
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(jwtSecret)

	authToken := &AuthToken{Token: token, ExpiresAt: expireTime.Format("2006-01-02 15:04:05")}
	return authToken, err
}

// ParseToken parsing token
func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}

	return nil, err
}
