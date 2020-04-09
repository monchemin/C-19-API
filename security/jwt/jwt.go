package jwt

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"os"
	"time"
)

type Claims struct {
	UserID string `json:"username"`
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
func GenerateToken(userID string) (tokenString string, err error) {
	expirationTime := time.Now().Add(tokenDuration * time.Minute)
	claims := &Claims{
		UserID: userID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
			IssuedAt: time.Now().Unix(),
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
func ParseToken(tokenString string)(claims *Claims, err error) {
	tkn, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil {
		return
	}
	if !tkn.Valid {
		err = errors.New("invalid token")
		return
	}
	return
}