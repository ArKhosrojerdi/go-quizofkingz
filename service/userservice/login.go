package userservice

import (
	"fmt"
	"quizofkingz/entity"

	"golang.org/x/crypto/bcrypt"
)

type LoginRequest struct {
	PhoneNumber string `json:"phone_number"`
	Password    string `json:"password"`
}

type LoginResponse struct {
	User entity.User `json:"user"`
}

func (s Service) Login(req LoginRequest) (LoginResponse, error) {
	user, exists, err := s.repo.GetUserByPhoneNumber(req.PhoneNumber)
	if err != nil {
		return LoginResponse{}, fmt.Errorf("phone number or password not correct")
	}

	if !exists {
		return LoginResponse{}, fmt.Errorf("phone number or password not correct")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		return LoginResponse{}, fmt.Errorf("phone number or password not correct")
	}

	return LoginResponse{}, nil
}
