package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"quizofkingz/repository/mysql"
	"quizofkingz/service/userservice"
)

func UserLoginHandler(writer http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodPost {
		fmt.Fprint(writer, "Method not allowed")
	}

	data, err := io.ReadAll(req.Body)
	if err != nil {
		writer.Write([]byte(fmt.Sprintf(`{"error": "%s"}`, err.Error())))
		return
	}

	var lReq userservice.LoginRequest
	err = json.Unmarshal(data, &lReq)
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
	_, err = userSvc.Login(lReq)
	if err != nil {
		writer.Write([]byte(fmt.Sprintf(`{"error": "%s"}`, err.Error())))
		return
	}
	writer.Write([]byte(`{"is_logged_in": true}`))
}
