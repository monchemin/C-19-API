package repository

import (
	"c19/errors"
	"c19/security/model"
)

func (r repository) CreateUser(request model.UserCreateRequest) (string, error) {
	if !request.IsValid() {
		return "", errors.InvalidRequestData()
	}
	row, err := r.db.NamedQuery(insertNewUser, request)
	if err != nil {
		return "", err
	}
	var ID string
	if row.Next() {
		row.Scan(&ID)
	}
	return ID, err
}

func (r repository) Login(request model.LoginRequest) LoginResult {
	if request.HasValidLogin() {
		return LoginResult{}
	}
	var result []LoginResult
	if err := r.db.Select(&result, userByEmail, request.Email, request.Password); err != nil {
		return LoginResult{}
	}
	return result[0]
}

func (r repository) StartSession(sessionID string) error {
	return nil
}

func (r repository) EndSession(sessionID string) {
	return
}

func (r repository) ChangePassword(userID, oldPassword, newPassword string) error {
	if len(userID) == 0 || len(newPassword) == 0 {
		return errors.InvalidRequestData()
	}
	var result []LoginResult
	if err := r.db.Select(&result, userByID, userID, oldPassword); err != nil {
		return err
	}
	if len(result) != 1 {
		return errors.InvalidRequestData()
	}
	_, err := r.db.Exec(changePassword, userID, newPassword)
	return err
}
