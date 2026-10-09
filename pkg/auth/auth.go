package auth

import (
	"crypto/rsa"
	"fmt"
	"os"
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
	tokenStr = strings.TrimSpace(strings.TrimPrefix(tokenStr, "Bearer "))
	if tokenStr == "" {
		return nil, fmt.Errorf("authorization token is empty")
	}

	privateKey, err := LoadPrivateKey()
	if err != nil {
		return nil, err
	}

	claims := &userservice.Claims{}
	token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (any, error) {
		if token.Method != jwt.SigningMethodRS256 {
			return nil, fmt.Errorf("unexpected JWT signing method: %s", token.Method.Alg())
		}
		return &privateKey.PublicKey, nil
	})
	if err != nil {
		return nil, err
	}
	if !token.Valid {
		return nil, fmt.Errorf("token is invalid")
	}
	return claims, nil
}
