package userservice

import (
	"crypto/rsa"
	"quizofkingz/entity"
	"quizofkingz/pkg/errmsg"
	"quizofkingz/pkg/richerror"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type ProfileRequest struct {
	UserID uint `json:"user_id"`
}

type ProfileResponse struct {
	User entity.User `json:"user"`
}

func (s Service) Profile(req ProfileRequest) (ProfileResponse, error) {
	const op richerror.Op = "userservice.Profile"

	user, err := s.repo.GetUserByID(req.UserID)
	if err != nil {
		return ProfileResponse{}, richerror.New(op).
			WithErr(err).
			WithKind(richerror.KindOf(err)).
			WithMessage(richerror.MessageOf(err))
	}

	return ProfileResponse{User: user}, nil
}

type CustomerInfo struct {
	Name string
	Kind string
}

type Claims struct {
	jwt.RegisteredClaims
	UserID    uint
	TokenType string
}

func (c Claims) Valid() error {
	return nil
}

const (
	AccessTokenType  = "access"
	RefreshTokenType = "refresh"
)

func (s Service) CreateTokenPair(userID uint) (LoginResponse, error) {
	const op richerror.Op = "userservice.CreateTokenPair"

	accessToken, err := createToken(userID, s.signKey, AccessTokenType, 2*time.Hour)
	if err != nil {
		return LoginResponse{}, richerror.New(op).
			WithErr(err).
			WithMessage(errmsg.ErrorMsgUnexpected).
			WithKind(richerror.KindUnexpected)
	}

	refreshToken, err := createToken(userID, s.signKey, RefreshTokenType, 30*24*time.Hour)
	if err != nil {
		return LoginResponse{}, richerror.New(op).
			WithErr(err).
			WithMessage(errmsg.ErrorMsgUnexpected).
			WithKind(richerror.KindUnexpected)
	}

	return LoginResponse{AccessToken: accessToken, RefreshToken: refreshToken}, nil
}

func createToken(userID uint, signKey *rsa.PrivateKey, tokenType string, ttl time.Duration) (string, error) {
	t := jwt.NewWithClaims(jwt.SigningMethodRS256, &Claims{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(ttl)),
		},
		UserID:    userID,
		TokenType: tokenType,
	})

	return t.SignedString(signKey)
}
