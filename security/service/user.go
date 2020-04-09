package service

import (
	"context"

	appContext "c19/context"
	"c19/errors"
	"c19/security/jwt"
	"c19/security/model"
)

func (s securityService) CreateUser(ctx context.Context, request model.UserCreateRequest) string {
	if ok := s.CanManage(ctx); !ok {
		return ""
	}
	if !request.IsValid() {
		return ""
	}
	return s.repository.CreateUser(request)
}

func (s securityService) Login(request model.LoginRequest) (model.LoginResponse, error) {
	if !request.HasValidLogin() {
		return model.LoginResponse{}, errors.InvalidRequestData()
	}

	result := s.repository.Login(request)
	if len(result.ID) == 0 {
		return model.LoginResponse{}, errors.InvalidRequestData()
	}

	tokenString, err := jwt.GenerateToken(result.ID)
	if err != nil {
		return model.LoginResponse{}, errors.InvalidRequestData()
	}

	if err = s.repository.StartSession(tokenString); err != nil {
		return model.LoginResponse{}, errors.InvalidRequestData()
	}

	return model.LoginResponse{Token: tokenString}, nil
}

func (s securityService) ChangePassword(ctx context.Context, request model.LoginRequest) error {
	if !request.HasValidPasswordChange() {
		return errors.InvalidRequestData()
	}

	ctxValues := appContext.ContextKeys(ctx)
	claims, err := jwt.ParseToken(ctxValues.Token)
	if err != nil {
		return err
	}

	return s.repository.ChangePassword(claims.UserID, request.NewPassword)
}

func (s securityService) Logout(ctx context.Context) {
	ctxValues := appContext.ContextKeys(ctx)
	s.repository.EndSession(ctxValues.Token)
}
