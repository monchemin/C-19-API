package repository

import "c19/security/model"

func (r repository) CreateUser(request model.UserCreateRequest) string {
	panic("implement me")
}

func (r repository) Login(request model.LoginRequest) LoginResult {
	panic("implement me")
}

func (r repository) StartSession(sessionID string) error {
	panic("implement me")
}

func (r repository) EndSession(sessionID string) {
	panic("implement me")
}

func (r repository) ChangePassword(userID, newPassword string) error {
	panic("implement me")
}
