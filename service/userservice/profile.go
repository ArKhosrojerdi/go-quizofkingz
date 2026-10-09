package userservice

import (
	"crypto/rsa"
	"fmt"
	"quizofkingz/entity"
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
	user, err := s.repo.GetUserByID(req.UserID)
	if err != nil {
		return ProfileResponse{}, fmt.Errorf("unexpected error: %w", err)
	}

	return ProfileResponse{User: user}, nil
}

type CustomerInfo struct {
	Name string
	Kind string
}

type Claims struct {
	jwt.RegisteredClaims
	UserID uint
}

func (c Claims) Valid() error {
	return nil
}

func createToken(userID uint, signKey *rsa.PrivateKey) (string, error) {
	t := jwt.New(jwt.SigningMethodRS256)

	t.Claims = &Claims{
		// set the expire time
		// see https://datatracker.ietf.org/doc/html/rfc7519#section-4.1.4
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24 * 7)),
		UserID:    userID,
	}

	return t.SignedString(signKey)
}
