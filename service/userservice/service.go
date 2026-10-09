package userservice

import (
	"crypto/rsa"
	"quizofkingz/entity"
)

type Repository interface {
	IsPhoneNumberUnique(phoneNumber string) (bool, error)
	RegisterUser(user entity.User) (entity.User, error)
	GetUserByPhoneNumber(phoneNumber string) (entity.User, bool, error)
	GetUserByID(userID uint) (entity.User, error)
}

type Service struct {
	signKey *rsa.PrivateKey
	repo    Repository
}

func New(repo Repository, signKey *rsa.PrivateKey) Service {
	return Service{repo: repo, signKey: signKey}
}
