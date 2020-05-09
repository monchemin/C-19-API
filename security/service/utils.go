package service

import (
	"context"
	"golang.org/x/crypto/bcrypt"

	appContext "github.com/monchemin/C-19-API/context"
	"github.com/monchemin/C-19-API/security/jwt"
)

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

func comparePassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func parseJwt(ctx context.Context) (*jwt.AppClaims, error) {
	ctxValues := appContext.ContextKeys(ctx)
	return jwt.ParseToken(ctxValues.Token)
}

