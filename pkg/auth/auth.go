package auth

import (
	"fmt"
	"log"
	"quizofkingz/service/userservice"
	"strings"

	"github.com/golang-jwt/jwt/v5"
)

func ParseJWT(tokenStr string) (*userservice.Claims, error) {
	tokenStr = strings.Replace(tokenStr, "Bearer ", "", 1)

	token, err := jwt.ParseWithClaims(tokenStr, &userservice.Claims{}, func(token *jwt.Token) (any, error) {
		return []byte("AllYourBase"), nil
	})

	if err != nil {
		log.Fatal(err)
	} else if claims, ok := token.Claims.(*userservice.Claims); ok {
		fmt.Printf("userID: %v, ExpiredAt:%v\n", claims.UserID, claims.RegisteredClaims.ExpiresAt)
		return claims, nil
	}

	return nil, err
}
