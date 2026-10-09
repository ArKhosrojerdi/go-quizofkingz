package httpserver

import (
	"net/http"
	"quizofkingz/pkg/auth"
	"quizofkingz/service/userservice"

	"github.com/labstack/echo/v5"
)

type RefreshRequest struct {
	RefreshToken string `json:"refresh_token"`
}

func (s Server) Login(c *echo.Context) error {
	var req userservice.LoginRequest
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid request payload: "+err.Error())
	}

	resp, err := s.userSvc.Login(req)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, err.Error())
	}

	return c.JSON(http.StatusOK, resp)
}

func (s Server) Register(c *echo.Context) error {
	var req userservice.RegisterRequest
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid request payload: "+err.Error())
	}

	resp, err := s.userSvc.Register(req)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusCreated, resp)
}

func (s Server) Profile(c *echo.Context) error {
	claims, err := auth.ParseJWT(c.Request().Header.Get("Authorization"))
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, "token is not valid")
	}
	if claims.TokenType != userservice.AccessTokenType {
		return echo.NewHTTPError(http.StatusUnauthorized, "token is not valid")
	}

	resp, err := s.userSvc.Profile(userservice.ProfileRequest{UserID: claims.UserID})
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, resp)
}

func (s Server) Refresh(c *echo.Context) error {
	var req RefreshRequest
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid request payload: "+err.Error())
	}

	claims, err := auth.ParseJWT(req.RefreshToken)
	if err != nil || claims.TokenType != userservice.RefreshTokenType {
		return echo.NewHTTPError(http.StatusUnauthorized, "refresh token is not valid")
	}

	resp, err := s.userSvc.CreateTokenPair(claims.UserID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, resp)
}
