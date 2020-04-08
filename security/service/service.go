package service

import (
	"c19/security/model"
	"c19/security/repository"
	"context"
)

type SecurityService interface {
	CreateUser(ctx context.Context, request model.UserCreateRequest) error
	CanManage(ctx context.Context, resourceID string) bool
	CanWrite(ctx context.Context, resourceID string) bool
	CanRead(ctx context.Context, resourceID string) bool
}

type securityService struct {
	repository repository.SecurityRepository
}

func NewSecurityService(repo repository.SecurityRepository) SecurityService {
	return &securityService{repository: repo}
}
