package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"quizofkingz/repository/mysql"
	"quizofkingz/service/userservice"
)

func GetProfile(writer http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodGet {
		fmt.Fprint(writer, "Method not allowed")
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

	userSvc := userservice.New(mysqlRepo)
	resp, err := userSvc.Profile(pReq)
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
