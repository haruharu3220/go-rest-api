package repository

import "go-rest-api/model"

type IUserRepository interface {
	GetUserEmail(user *model.User, email string) error
	CreateUser(user *model.User) error
}