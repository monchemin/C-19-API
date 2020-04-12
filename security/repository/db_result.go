package repository

import "github.com/google/uuid"

type LoginResult struct {
	ID        string `db:"id"`
	Email     string `db:"email"`
	Password  string `db:"password"`
	FirstName string `db:"first_name"`
	LastName  string `db:"last_name"`
	IsActive  bool   `db:"active"`
}

type PrivilegeResult struct {
	Code string `db:"code"`
	Name string `db:"name"`
}

type PrivilegeRequest struct {
	RoleID         int       `json:"role_id"`
	ResourceTypeID int       `json:"resource_type_id"`
	ResourceID     uuid.UUID `json:"resource_id"`
	UserID         uuid.UUID `json:"user_id"`
}
