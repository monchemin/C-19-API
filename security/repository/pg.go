package repository

import (
	"c19/errors"
	"c19/security/model"
	"github.com/google/uuid"
	"log"
)

func (r repository) CreateUser(request model.UserCreateRequest) (string, error) {
	if !request.IsValid() {
		return "", errors.InvalidRequestData()
	}
	row, err := r.db.NamedQuery(insertNewUser, request)
	if err != nil {
		return "", err
	}
	var ID uuid.UUID
	if row.Next() {
		row.Scan(&ID)
	}
	pv := make([]PrivilegeRequest, len(request.Privileges))
	for i, p := range request.Privileges {
		pv[i] = PrivilegeRequest{
			RoleID:         p.RoleID,
			ResourceTypeID: p.ResourceTypeID,
			ResourceID:     uuid.MustParse(p.ResourceID),
			UserID:         ID,
		}
	}

	_, err = r.db.NamedQuery(insertPrivilege, pv)
	log.Println(err)
	return ID.String(), err
}

func (r repository) Login(request model.LoginRequest) LoginResult {
	if !request.HasValidLogin() {
		return LoginResult{}
	}
	var result []LoginResult
	if err := r.db.Select(&result, userByEmail, request.Email); err != nil {
		log.Println(err)
		return LoginResult{}
	}
	if len(result) == 0 {
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

func (r repository) ChangePassword(userID, newPassword string) error {
	if len(userID) == 0 || len(newPassword) == 0 {
		return errors.InvalidRequestData()
	}
	_, err := r.db.Exec(changePassword, userID, newPassword)
	return err
}

func (r repository) UserByID(userID string) LoginResult {
	var result []LoginResult
	if err := r.db.Select(&result, userByID, userID); err != nil {
		log.Println(err)
		return LoginResult{}
	}
	if len(result) == 0 {
		return LoginResult{}
	}
	return result[0]
}

func (r repository) UserPrivileges(userID, resourceID string) []PrivilegeResult {
	var result []PrivilegeResult
	if err := r.db.Select(&result, userPrivileges, userID, resourceID); err != nil {
		log.Println(err)
		return nil
	}
	return result
}
