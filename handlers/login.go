package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"quizofkingz/repository/mysql"
	"quizofkingz/service/userservice"

	"github.com/golang-jwt/jwt/v5"
)

func UserLoginHandler(writer http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodPost {
		fmt.Fprint(writer, "Method not allowed")
	}

	data, err := io.ReadAll(req.Body)
	if err != nil {
		writeLoginError(writer, http.StatusBadRequest, err.Error())
		return
	}

	var lReq userservice.LoginRequest
	err = json.Unmarshal(data, &lReq)
	if err != nil {
		writeLoginError(writer, http.StatusBadRequest, err.Error())
		return
	}

	mysqlRepo, err := mysql.New()

	if err != nil {
		writeLoginError(writer, http.StatusInternalServerError, err.Error())
		return
	}

	keyPath := os.Getenv("JWT_PRIVATE_KEY_PATH")
	if keyPath == "" {
		writeLoginError(writer, http.StatusInternalServerError, "JWT_PRIVATE_KEY_PATH is not set")
		return
	}

	keyPEM, err := os.ReadFile(keyPath)
	if err != nil {
		writeLoginError(writer, http.StatusInternalServerError, fmt.Sprintf("read JWT private key: %s", err))
		return
	}

	privateKey, err := jwt.ParseRSAPrivateKeyFromPEM(keyPEM)
	if err != nil {
		writeLoginError(writer, http.StatusInternalServerError, fmt.Sprintf("parse JWT private key: %s", err))
		return
	}

	userSvc := userservice.New(mysqlRepo, privateKey)
	resp, err := userSvc.Login(lReq)
	if err != nil {
		writeLoginError(writer, http.StatusUnauthorized, err.Error())
		return
	}

	data, err = json.Marshal(resp)
	if err != nil {
		writeLoginError(writer, http.StatusInternalServerError, err.Error())
		return
	}

	writer.Write(data)
}

func writeLoginError(writer http.ResponseWriter, status int, message string) {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(status)
	_ = json.NewEncoder(writer).Encode(map[string]string{"error": message})
}
