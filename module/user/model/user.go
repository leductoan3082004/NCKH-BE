package usermodel

import (
	"errors"
	"nckh-BE/appCommon"
	"net/http"
)

const EntityName = "User"

type User struct {
	appCommon.MgDBModel `json:",inline"`
	Username            string `json:"username" bson:"username"`
	Password            string `json:"password" bson:"password"`
	Salt                string `json:"salt" bson:"salt"`
	Name                string `json:"name" bson:"name"`
	IsAdmin             bool   `json:"is_admin" bson:"is_admin"`
}

type UserRegister struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Name     string `json:"name" binding:"required"`
}
type UserLogin struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (User) TableName() string {
	return "user"
}

var (
	ErrUsernameExisted = appCommon.NewCustomError(
		http.StatusBadRequest,
		errors.New("username has already existed"),
		"username has already existed",
		"ErrUsernameExisted",
	)
)
