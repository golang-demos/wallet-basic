package models

import "github.com/go-playground/validator/v10"

type UserShort struct {
	ID   string `json:"id"`
	Name string `json:"name" validate:"required,min=2,max=32"`
	Role string `json:"role" validate:"required,min=2,max=16"`
}
type User struct {
	UserShort
	Token string `json:"token" validate:"required,min=24,max=24"`
}

func (u *User) ToShort() *UserShort {
	var shortUser UserShort
	shortUser.ID = u.ID
	shortUser.Name = u.Name
	shortUser.Role = u.Role
	return &shortUser
}

func (u *User) Validate() []*ErrorResp {
	return validateModelsForErrors(validator.New().Struct(u))
}
