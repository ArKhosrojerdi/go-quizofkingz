package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"quizofkingz/repository/mysql"
	"quizofkingz/service/userservice"
)

func UserRegisterHandler(writer http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodPost {
		fmt.Fprint(writer, "Method not allowed")
	}

	data, err := io.ReadAll(req.Body)
	if err != nil {
		writer.Write([]byte(fmt.Sprintf(`{"error": "%s"}`, err.Error())))
		return
	}

	var uReq userservice.RegisterRequest
	err = json.Unmarshal(data, &uReq)
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
	_, err = userSvc.Register(uReq)
	if err != nil {
		writer.Write([]byte(fmt.Sprintf(`{"error": "%s"}`, err.Error())))
		return
	}
	writer.Write([]byte(`{"message": "user created"}`))
}
