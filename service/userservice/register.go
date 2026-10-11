package userservice

import (
	"errors"
	"quizofkingz/entity"
	"quizofkingz/pkg/errmsg"
	"quizofkingz/pkg/phonenumber"
	"quizofkingz/pkg/richerror"

	validation "github.com/go-ozzo/ozzo-validation/v4"
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

func (r RegisterRequest) Validate() error {
	return validation.ValidateStruct(&r,
		validation.Field(&r.Name,
			validation.Required.Error(errmsg.ErrorMsgUsernameRequired),
			validation.Length(3, 50).Error(errmsg.ErrorMsgUsernameLength),
		),
		validation.Field(&r.PhoneNumber,
			validation.Required.Error(errmsg.ErrorMsgPhoneNumberRequired),
			validation.By(validatePhoneNumber),
		),
		validation.Field(&r.Password,
			validation.Required.Error(errmsg.ErrorMsgPasswordRequired),
			validation.Length(4, 72).Error(errmsg.ErrorMsgPasswordLength),
		),
	)
}

func validatePhoneNumber(value any) error {
	phone, _ := value.(string)
	if !phonenumber.IsValid(phone) {
		return errors.New(errmsg.ErrorMsgPhoneNumberIsNotValid)
	}
	return nil
}

func (s Service) Register(req RegisterRequest) (RegisterResponse, error) {
	const op richerror.Op = "userservice.Register"
	// TODO - add OTP

	if err := req.Validate(); err != nil {
		var fields validation.Errors
		if errors.As(err, &fields) {
			return RegisterResponse{}, richerror.New(op).
				WithErr(fields).
				WithMessage(errmsg.ErrorMsgInvalidFields).
				WithKind(richerror.KindInvalid)
		}

		return RegisterResponse{}, richerror.New(op).
			WithErr(err).
			WithMessage(errmsg.ErrorMsgUnexpected).
			WithKind(richerror.KindUnexpected)
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
