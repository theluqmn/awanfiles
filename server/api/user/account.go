package user

import (
	"net/http"
	"github.com/labstack/echo"
)

func CreateAccount(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}