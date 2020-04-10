package service

import (
	"context"
	"golang.org/x/crypto/bcrypt"

	appContext "c19/context"
	"c19/errors"
	"c19/security/jwt"
	"c19/security/model"
)

func (ss securityService) CreateUser(ctx context.Context, request model.UserCreateRequest) (string, error) {
	if ok := ss.CanManage(ctx); !ok {
		return "", errors.Unauthorized()
	}
	if !request.IsValid() {
		return "", errors.InvalidRequestData()
	}
	claims, err := parseJwt(ctx)
	if err != nil {
		return "", err
	}
	hashPass, _ := hashPassword(request.Password)
	request.Password = hashPass
	request.CreatedBy = claims.UserID
	return ss.repository.CreateUser(request)
}

func (ss securityService) Login(request model.LoginRequest) (model.LoginResponse, error) {
	if !request.HasValidLogin() {
		return model.LoginResponse{}, errors.InvalidRequestData()
	}
	hashPass, _ := hashPassword(request.Password)
	request.Password = hashPass
	result := ss.repository.Login(request)
	if len(result.ID) == 0 {
		return model.LoginResponse{}, errors.InvalidRequestData()
	}

	tokenString, err := jwt.GenerateToken(result.ID)
	if err != nil {
		return model.LoginResponse{}, errors.InvalidRequestData()
	}

	if err = ss.repository.StartSession(tokenString); err != nil {
		return model.LoginResponse{}, errors.InvalidRequestData()
	}

	return model.LoginResponse{Token: tokenString}, nil
}

func (ss securityService) ChangePassword(ctx context.Context, request model.LoginRequest) error {
	if !request.HasValidPasswordChange() {
		return errors.InvalidRequestData()
	}

	claims, err := parseJwt(ctx)
	if err != nil {
		return err
	}
	hashOldPass, _ := hashPassword(request.Password)
	hashNewPass, _ := hashPassword(request.NewPassword)
	return ss.repository.ChangePassword(claims.UserID, hashOldPass, hashNewPass)
}

func (ss securityService) Logout(ctx context.Context) {
	ctxValues := appContext.ContextKeys(ctx)
	ss.repository.EndSession(ctxValues.Token)
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func parseJwt(ctx context.Context) (*jwt.Claims, error) {
	ctxValues := appContext.ContextKeys(ctx)
	return jwt.ParseToken(ctxValues.Token)
}