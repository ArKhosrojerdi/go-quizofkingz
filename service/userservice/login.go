package userservice

import (
	"quizofkingz/pkg/errmsg"
	"quizofkingz/pkg/richerror"

	"golang.org/x/crypto/bcrypt"
)

type LoginRequest struct {
	PhoneNumber string `json:"phone_number"`
	Password    string `json:"password"`
}

type LoginResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

func (s Service) Login(req LoginRequest) (LoginResponse, error) {
	const op richerror.Op = "userservice.Login"

	user, exists, err := s.repo.GetUserByPhoneNumber(req.PhoneNumber)
	if err != nil {
		return LoginResponse{}, richerror.New(op).
			WithErr(err).
			WithMessage(errmsg.ErrorMsgUnexpected).
			WithKind(richerror.KindUnexpected)
	}

	if !exists {
		return LoginResponse{}, richerror.New(op).
			WithMessage(errmsg.ErrorMsgPhoneNumberOrPasswordIncorrect).
			WithKind(richerror.KindUnauthorized)
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		return LoginResponse{}, richerror.New(op).
			WithErr(err).
			WithMessage(errmsg.ErrorMsgPhoneNumberOrPasswordIncorrect).
			WithKind(richerror.KindUnauthorized)
	}

	tokens, err := s.CreateTokenPair(user.ID)
	if err != nil {
		return LoginResponse{}, richerror.New(op).
			WithErr(err).
			WithMessage(errmsg.ErrorMsgUnexpected).
			WithKind(richerror.KindUnexpected)
	}

	return tokens, nil
}
