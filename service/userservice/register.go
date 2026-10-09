package userservice

import (
	"fmt"
	"quizofkingz/entity"
	"quizofkingz/pkg/phonenumber"

	"golang.org/x/crypto/bcrypt"
)

type RegisterRequest struct {
	Name        string `json:"username"`
	PhoneNumber string `json:"phone_number"`
	Password    string `json:"password"`
}

type RegisterResponse struct {
	User entity.User `json:"user"`
}

func (s Service) Register(req RegisterRequest) (RegisterResponse, error) {
	// TODO - add OTP

	if !phonenumber.IsValid(req.PhoneNumber) {
		return RegisterResponse{}, fmt.Errorf("Phone number is not valid")
	}

	if isUnique, err := s.repo.IsPhoneNumberUnique(req.PhoneNumber); err != nil || !isUnique {
		if err != nil {
			return RegisterResponse{}, err
		}

		return RegisterResponse{}, fmt.Errorf("Phone number is not unique")
	}

	if len(req.Name) < 3 {
		return RegisterResponse{}, fmt.Errorf("Name is too short")
	}

	if len(req.Password) < 4 {
		return RegisterResponse{}, fmt.Errorf("Password is too short")
	}

	pass, err := bcrypt.GenerateFromPassword([]byte(req.Password), 0)
	if err != nil {
		return RegisterResponse{}, fmt.Errorf("Failed to create user")
	}

	user := entity.User{
		ID:          0,
		Name:        req.Name,
		PhoneNumber: req.PhoneNumber,
		Password:    string(pass),
	}

	createdUser, err := s.repo.RegisterUser(user)
	if err != nil {
		return RegisterResponse{}, fmt.Errorf("unexpected error: %w", err)
	}

	return RegisterResponse{User: createdUser}, nil
}
