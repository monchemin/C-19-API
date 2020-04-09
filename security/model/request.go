package model

import "regexp"

type UserCreateRequest struct {
	Email       string `json:"e_mail"`
	Password    string `json:"password"`
	PhoneNumber string `json:"phone_number"`
	FirstName   string `json:"phone_number"`
	LastName    string `json:"phone_number"`
}

var rxEmail = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

func (us *UserCreateRequest) IsValid() bool {
	if !rxEmail.MatchString(us.Email) {
		return false
	}
	if len(us.Password) < 5 {
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
	return rxEmail.MatchString(us.Email) && us.Password != ""
}

func (us *LoginRequest) HasValidPasswordChange() bool {
	return us.NewPassword != "" && us.Password != ""
}
