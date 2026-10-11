package httpserver

import (
	"log"
	"net/http"
	"quizofkingz/pkg/auth"
	"quizofkingz/pkg/errmsg"
	"quizofkingz/pkg/richerror"
	"quizofkingz/service/userservice"

	"github.com/labstack/echo/v5"
)

type RefreshRequest struct {
	RefreshToken string `json:"refresh_token"`
}

func (s Server) Login(c *echo.Context) error {
	var req userservice.LoginRequest
	if err := c.Bind(&req); err != nil {
		log.Printf("bind login: %v", err)
		return echo.NewHTTPError(http.StatusBadRequest, errmsg.ErrorMsgInvalidRequest)
	}

	resp, err := s.userSvc.Login(req)
	if err != nil {
		return toHTTPError(err)
	}

	return c.JSON(http.StatusOK, resp)
}

func (s Server) Register(c *echo.Context) error {
	var req userservice.RegisterRequest
	if err := c.Bind(&req); err != nil {
		log.Printf("bind register: %v", err)
		return echo.NewHTTPError(http.StatusBadRequest, errmsg.ErrorMsgInvalidRequest)
	}

	resp, err := s.userSvc.Register(req)
	if err != nil {
		return writeServiceError(c, err)
	}

	return c.JSON(http.StatusCreated, resp)
}

func (s Server) Profile(c *echo.Context) error {
	claims, err := auth.ParseJWT(c.Request().Header.Get("Authorization"))
	if err != nil {
		return toHTTPError(err)
	}
	if claims.TokenType != userservice.AccessTokenType {
		return echo.NewHTTPError(http.StatusUnauthorized, errmsg.ErrorMsgTokenNotValid)
	}

	resp, err := s.userSvc.Profile(userservice.ProfileRequest{UserID: claims.UserID})
	if err != nil {
		return toHTTPError(err)
	}

	return c.JSON(http.StatusOK, resp)
}

func (s Server) Refresh(c *echo.Context) error {
	var req RefreshRequest
	if err := c.Bind(&req); err != nil {
		log.Printf("bind refresh: %v", err)
		return echo.NewHTTPError(http.StatusBadRequest, errmsg.ErrorMsgInvalidRequest)
	}

	claims, err := auth.ParseJWT(req.RefreshToken)
	if err != nil {
		if richerror.KindOf(err) == richerror.KindUnexpected {
			return toHTTPError(err)
		}
		return echo.NewHTTPError(http.StatusUnauthorized, errmsg.ErrorMsgRefreshTokenNotValid)
	}
	if claims.TokenType != userservice.RefreshTokenType {
		return echo.NewHTTPError(http.StatusUnauthorized, errmsg.ErrorMsgRefreshTokenNotValid)
	}

	resp, err := s.userSvc.CreateTokenPair(claims.UserID)
	if err != nil {
		return toHTTPError(err)
	}

	return c.JSON(http.StatusOK, resp)
}
