package userservice

import (
	"crypto/rand"
	"crypto/rsa"
	"testing"

	"github.com/golang-jwt/jwt/v5"
)

func TestCreateTokenSignsWithRS256(t *testing.T) {
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		t.Fatal(err)
	}

	tokenString, err := createToken(42, privateKey)
	if err != nil {
		t.Fatal(err)
	}

	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		if token.Method != jwt.SigningMethodRS256 {
			t.Fatalf("unexpected signing method: %v", token.Method.Alg())
		}
		return &privateKey.PublicKey, nil
	})
	if err != nil {
		t.Fatal(err)
	}
	if !token.Valid {
		t.Fatal("token is invalid")
	}
	if claims.UserID != 42 {
		t.Fatalf("got user ID %d, want 42", claims.UserID)
	}
}
