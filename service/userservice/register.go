package userservice

import (
	"quizofkingz/entity"
	"quizofkingz/pkg/errmsg"
	"quizofkingz/pkg/phonenumber"
	"quizofkingz/pkg/richerror"

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
	const op richerror.Op = "userservice.Register"
	// TODO - add OTP

	if !phonenumber.IsValid(req.PhoneNumber) {
		return RegisterResponse{}, richerror.New(op).
			WithMessage(errmsg.ErrorMsgPhoneNumberIsNotValid).
			WithKind(richerror.KindInvalid)
	}

	isUnique, err := s.repo.IsPhoneNumberUnique(req.PhoneNumber)
	if err != nil {
		return RegisterResponse{}, richerror.New(op).
			WithErr(err).
			WithMessage(errmsg.ErrorMsgUnexpected).
			WithKind(richerror.KindUnexpected)
	}
	if !isUnique {
		return RegisterResponse{}, richerror.New(op).
			WithMessage(errmsg.ErrorMsgPhoneNumberIsNotUnique).
			WithKind(richerror.KindConflict)
	}

	if len(req.Name) < 3 {
		return RegisterResponse{}, richerror.New(op).
			WithMessage(errmsg.ErrorMsgNameTooShort).
			WithKind(richerror.KindInvalid)
	}

	if len(req.Password) < 4 {
		return RegisterResponse{}, richerror.New(op).
			WithMessage(errmsg.ErrorMsgPasswordTooShort).
			WithKind(richerror.KindInvalid)
	}

	pass, err := bcrypt.GenerateFromPassword([]byte(req.Password), 0)
	if err != nil {
		return RegisterResponse{}, richerror.New(op).
			WithErr(err).
			WithMessage(errmsg.ErrorMsgUnexpected).
			WithKind(richerror.KindUnexpected)
	}

	user := entity.User{
		ID:          0,
		Name:        req.Name,
		PhoneNumber: req.PhoneNumber,
		Password:    string(pass),
	}

	createdUser, err := s.repo.RegisterUser(user)
	if err != nil {
		return RegisterResponse{}, richerror.New(op).
			WithErr(err).
			WithMessage(errmsg.ErrorMsgUnexpected).
			WithKind(richerror.KindUnexpected)
	}

	return RegisterResponse{User: createdUser}, nil
}
