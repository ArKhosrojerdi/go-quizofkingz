package userservice

import (
	"quizofkingz/entity"
)

type Repository interface {
	IsPhoneNumberUnique(phoneNumber string) (bool, error)
	RegisterUser(user entity.User) (entity.User, error)
	GetUserByPhoneNumber(phoneNumber string) (entity.User, bool, error)
	GetUserByID(userID uint) (entity.User, error)
}

type Service struct {
	repo Repository
}

func New(repo Repository) Service {
	return Service{repo: repo}
}
