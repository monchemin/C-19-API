package model

import (
	"os"
	"regexp"
)

type UserCreateRequest struct {
	Email       string      `json:"e_mail"`
	Password    string      `json:"password"`
	PhoneNumber string      `json:"phone_number"`
	FirstName   string      `json:"first_name"`
	LastName    string      `json:"last_name"`
	CreatedBy   string      `json:"created_by"`
	ResourceID  string      `json:"resource_id"`
	Privileges  []Privilege `json:"privileges"`
}

type Privilege struct {
	RoleID         int    `json:"role_id"`
	ResourceTypeID int    `json:"resource_type_id"`
	ResourceID     string `json:"resource_id"`
}

var rxEmail = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

func (us *UserCreateRequest) IsValid() bool {
	//if us.Email == os.Getenv("SU") {
	//	return true
	//}
	if !rxEmail.MatchString(us.Email) {
		return false
	}
	if len(us.Password) < 5 {
		return false
	}
	if len(us.FirstName) == 0 || len(us.LastName) == 0 {
		return false
	}
	return true
}

type LoginRequest struct {
	Email       string `json:"e_mail"`
	Password    string `json:"password"`
	NewPassword string `json:"new_password"`
}

func (us *LoginRequest) HasValidLogin() bool {
	if us.Email == os.Getenv("SU") {
		return true
	}
	return rxEmail.MatchString(us.Email) && us.Password != ""
}

func (us *LoginRequest) HasValidPasswordChange() bool {
	return us.NewPassword != "" && us.Password != ""
}
