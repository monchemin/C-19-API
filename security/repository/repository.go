package repository

import (
	"github.com/monchemin/C-19-API/connector/pgsql"
	"github.com/monchemin/C-19-API/security/model"
)


type SecurityRepository interface {
	CreateUser(request model.UserCreateRequest) (string, error)
	Login(request model.LoginRequest) LoginResult
	StartSession(sessionID string) error
	ChangePassword(userID, newPassword string) error
	EndSession(sessionID string)
	UserByID(userID string) LoginResult
	UserPrivileges(userID, resourceID string) []PrivilegeResult
}

type repository struct {
	db *pgsql.DB
}

func NewSecurityRepository(db *pgsql.DB) SecurityRepository {
	return repository{db: db}
}