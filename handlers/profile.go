package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"quizofkingz/pkg/auth"
	"quizofkingz/repository/mysql"
	"quizofkingz/service/userservice"
)

func GetProfile(writer http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodGet {
		fmt.Fprint(writer, "Method not allowed")
	}

	authen := req.Header.Get("Authorization")
	claims, err := auth.ParseJWT(authen)
	if err != nil {
		fmt.Fprint(writer, http.StatusUnauthorized, `{"error": "token is not valid"}`)
	}

	data, err := io.ReadAll(req.Body)
	if err != nil {
		writer.Write([]byte(fmt.Sprintf(`{"error": "%s"}`, err.Error())))
		return
	}

	var pReq userservice.ProfileRequest
	err = json.Unmarshal(data, &pReq)
	if err != nil {
		writer.Write([]byte(fmt.Sprintf(`{"error": "%s"}`, err.Error())))
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

	data, err = json.Marshal(resp)
	if err != nil {
		writer.Write([]byte(fmt.Sprintf(`{"error": "%s"}`, err.Error())))
		return
	}

	writer.Write(data)
}
