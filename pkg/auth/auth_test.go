package auth_test

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"os"
	"path/filepath"
	"testing"
	"time"

	"quizofkingz/pkg/auth"
	"quizofkingz/service/userservice"

	"github.com/golang-jwt/jwt/v5"
)

func TestParseJWTVerifiesRS256WithConfiguredPrivateKey(t *testing.T) {
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		t.Fatal(err)
	}

	keyPath := filepath.Join(t.TempDir(), "jwt-private-key.pem")
	keyPEM := pem.EncodeToMemory(&pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(privateKey),
	})
	if err := os.WriteFile(keyPath, keyPEM, 0600); err != nil {
		t.Fatal(err)
	}
	t.Setenv("JWT_PRIVATE_KEY_PATH", keyPath)

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, &userservice.Claims{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour)),
		},
		UserID: 42,
	})
	tokenString, err := token.SignedString(privateKey)
	if err != nil {
		t.Fatal(err)
	}

	claims, err := auth.ParseJWT("Bearer " + tokenString)
	if err != nil {
		t.Fatal(err)
	}
	if claims.UserID != 42 {
		t.Fatalf("got user ID %d, want 42", claims.UserID)
	}
}
