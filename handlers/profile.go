package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"quizofkingz/pkg/auth"
	"quizofkingz/repository/mysql"
	"quizofkingz/service/userservice"
)

func GetProfile(writer http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodGet {
		writer.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprint(writer, "Method not allowed")
		return
	}

	authen := req.Header.Get("Authorization")
	claims, err := auth.ParseJWT(authen)
	if err != nil {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusUnauthorized)
		fmt.Fprint(writer, `{"error": "token is not valid"}`)
		return
	}

	mysqlRepo, err := mysql.New()

	if err != nil {
		writer.Write([]byte(fmt.Sprintf(`{"error": "%s"}`, err.Error())))
		return
	}

	userSvc := userservice.New(mysqlRepo, nil)
	resp, err := userSvc.Profile(userservice.ProfileRequest{
		UserID: claims.UserID,
	})
	if err != nil {
		writer.Write([]byte(fmt.Sprintf(`{"error": "%s"}`, err.Error())))
		return
	}

	data, err := json.Marshal(resp)
	if err != nil {
		writer.Write([]byte(fmt.Sprintf(`{"error": "%s"}`, err.Error())))
		return
	}

	writer.Write(data)
}
