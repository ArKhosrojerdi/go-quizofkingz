package httpserver

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo/v5"
	"quizofkingz/service/userservice"
)

func TestRegisterFieldErrorBody(t *testing.T) {
	srv := Server{userSvc: userservice.New(nil, nil)}
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/users/register", strings.NewReader(`{"username":"ab","phone_number":"123","password":"1234"}`))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	if err := srv.Register(c); err != nil {
		t.Fatal(err)
	}
	if rec.Code != http.StatusBadRequest {
		t.Fatalf("status = %d, body = %s", rec.Code, rec.Body.String())
	}

	var body struct {
		Error   map[string]string `json:"error"`
		Message string            `json:"message"`
	}
	if err := json.Unmarshal(rec.Body.Bytes(), &body); err != nil {
		t.Fatal(err)
	}
	if body.Message != "invalid fields" {
		t.Fatalf("message = %q", body.Message)
	}
	if body.Error["username"] != "username must be between 3 and 50 characters" {
		t.Fatalf("username = %q", body.Error["username"])
	}
	if body.Error["phone_number"] != "phone number is not valid" {
		t.Fatalf("phone_number = %q", body.Error["phone_number"])
	}
	if _, ok := body.Error["password"]; ok {
		t.Fatal("password should be valid")
	}
}
