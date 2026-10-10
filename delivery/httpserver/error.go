package httpserver

import (
	"log"
	"net/http"
	"quizofkingz/pkg/errmsg"
	"quizofkingz/pkg/richerror"

	"github.com/labstack/echo/v5"
)

func toHTTPError(err error) *echo.HTTPError {
	kind := richerror.KindOf(err)
	message := richerror.MessageOf(err)
	if kind == richerror.KindUnexpected || message == "" {
		log.Printf("%v", err)
		message = errmsg.ErrorMsgUnexpected
	}

	return echo.NewHTTPError(statusFromKind(kind), message)
}

func statusFromKind(kind richerror.Kind) int {
	switch kind {
	case richerror.KindInvalid:
		return http.StatusBadRequest
	case richerror.KindUnauthorized:
		return http.StatusUnauthorized
	case richerror.KindNotFound:
		return http.StatusNotFound
	case richerror.KindConflict:
		return http.StatusConflict
	default:
		return http.StatusInternalServerError
	}
}
