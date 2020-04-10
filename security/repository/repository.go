package repository

import (
	"c19/connector/pgsql"
	"c19/security/model"
)


type SecurityRepository interface {
	CreateUser(request model.UserCreateRequest) (string, error)
	Login(request model.LoginRequest) LoginResult
	StartSession(sessionID string) error
	ChangePassword(userID, oldPassword, newPassword string) error
	EndSession(sessionID string)
}

type repository struct {
	db *pgsql.DB
}

func NewSecurityRepository(db *pgsql.DB) SecurityRepository {
	return repository{db: db}
}