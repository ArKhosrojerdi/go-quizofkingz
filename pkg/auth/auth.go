package auth

import (
	"crypto/rsa"
	"fmt"
	"os"
	"quizofkingz/pkg/errmsg"
	"quizofkingz/pkg/richerror"
	"quizofkingz/service/userservice"
	"strings"

	"github.com/golang-jwt/jwt/v5"
)

func LoadPrivateKey() (*rsa.PrivateKey, error) {
	keyPath := os.Getenv("JWT_PRIVATE_KEY_PATH")
	if keyPath == "" {
		return nil, fmt.Errorf("JWT_PRIVATE_KEY_PATH is not set")
	}

	keyPEM, err := os.ReadFile(keyPath)
	if err != nil {
		return nil, fmt.Errorf("read JWT private key: %w", err)
	}

	privateKey, err := jwt.ParseRSAPrivateKeyFromPEM(keyPEM)
	if err != nil {
		return nil, fmt.Errorf("parse JWT private key: %w", err)
	}
	return privateKey, nil
}

func ParseJWT(tokenStr string) (*userservice.Claims, error) {
	const op richerror.Op = "auth.ParseJWT"

	tokenStr = strings.TrimSpace(strings.TrimPrefix(tokenStr, "Bearer "))
	if tokenStr == "" {
		return nil, richerror.New(op).
			WithMessage(errmsg.ErrorMsgTokenNotValid).
			WithKind(richerror.KindUnauthorized)
	}

	privateKey, err := LoadPrivateKey()
	if err != nil {
		return nil, richerror.New(op).
			WithErr(err).
			WithMessage(errmsg.ErrorMsgUnexpected).
			WithKind(richerror.KindUnexpected)
	}

	claims := &userservice.Claims{}
	token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (any, error) {
		if token.Method != jwt.SigningMethodRS256 {
			return nil, fmt.Errorf("unexpected JWT signing method: %s", token.Method.Alg())
		}
		return &privateKey.PublicKey, nil
	})
	if err != nil || !token.Valid {
		return nil, richerror.New(op).
			WithErr(err).
			WithMessage(errmsg.ErrorMsgTokenNotValid).
			WithKind(richerror.KindUnauthorized)
	}
	return claims, nil
}
