package jwt

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"os"
	"time"
)

type AppClaims struct {
	UserID   string `json:"user_id"`
	UserName string `json:"user_name"`
	jwt.StandardClaims
}

const (
	tokenDuration = 5
	expireOffset  = 3600
)

var jwtKey = []byte(os.Getenv("TOKEN_KEY"))

// input userID as string
// output signed token and and error
// purpose: encrypt user information
func GenerateToken(userID, userName string) (tokenString string, err error) {
	expirationTime := time.Now().Add(tokenDuration * time.Hour)
	claims := &AppClaims{
		UserID:   userID,
		UserName: userName,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	}
	// Declare the token with the algorithm used for signing, and the claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// Create the JWT string
	tokenString, err = token.SignedString(jwtKey)
	return
}

// input token string
// output claims and error
// if parsing is ok (not expire and signature match) error will be nil
func ParseToken(tokenString string) (*AppClaims, error) {
	claims := &AppClaims{}
	tkn, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil {
		return nil, err
	}
	if !tkn.Valid {
		return nil, errors.New("invalid token")
	}
	return claims, nil
}
