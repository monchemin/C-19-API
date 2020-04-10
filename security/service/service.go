package service

import (
	"c19/security/model"
	"c19/security/repository"
	"context"
)

type SecurityService interface {
	CreateUser(ctx context.Context, request model.UserCreateRequest) (string, error)
	Login(request model.LoginRequest) (model.LoginResponse, error)
	ChangePassword(ctx context.Context, request model.LoginRequest) error
	Logout(ctx context.Context)

	CanManage(ctx context.Context) bool
	CanWrite(ctx context.Context) bool
	CanRead(ctx context.Context) bool
}

type securityService struct {
	repository repository.SecurityRepository
}

func NewSecurityService(repo repository.SecurityRepository) SecurityService {
	return &securityService{repository: repo}
}
