package userservice

import (
	"crypto/rand"
	"crypto/rsa"
	"testing"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func TestCreateTokenSignsWithRS256(t *testing.T) {
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		t.Fatal(err)
	}

	tokenString, err := createToken(42, privateKey, AccessTokenType, 2*time.Hour)
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
	if claims.TokenType != AccessTokenType {
		t.Fatalf("got token type %q, want %q", claims.TokenType, AccessTokenType)
	}
}

func TestCreateTokenPairUsesExpectedLifetimes(t *testing.T) {
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		t.Fatal(err)
	}

	before := time.Now()
	tokens, err := New(privateKey, nil).CreateTokenPair(42)
	if err != nil {
		t.Fatal(err)
	}

	for _, tc := range []struct {
		name      string
		token     string
		tokenType string
		ttl       time.Duration
	}{
		{name: "access", token: tokens.AccessToken, tokenType: AccessTokenType, ttl: 2 * time.Hour},
		{name: "refresh", token: tokens.RefreshToken, tokenType: RefreshTokenType, ttl: 30 * 24 * time.Hour},
	} {
		t.Run(tc.name, func(t *testing.T) {
			claims := &Claims{}
			token, err := jwt.ParseWithClaims(tc.token, claims, func(token *jwt.Token) (interface{}, error) {
				return &privateKey.PublicKey, nil
			})
			if err != nil {
				t.Fatal(err)
			}
			if !token.Valid {
				t.Fatal("token is invalid")
			}
			if claims.TokenType != tc.tokenType {
				t.Fatalf("got token type %q, want %q", claims.TokenType, tc.tokenType)
			}
			if claims.ExpiresAt == nil {
				t.Fatal("token has no expiry")
			}
			expiry := claims.ExpiresAt.Time
			if expiry.Before(before.Add(tc.ttl-time.Second)) || expiry.After(time.Now().Add(tc.ttl)) {
				t.Fatalf("token expires at %v, want roughly %v from now", expiry, tc.ttl)
			}
		})
	}
}
