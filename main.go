package main

import (
	"main/utils"

	"net/http"
	"github.com/labstack/echo"
)

func main() {
	server := echo.New()
	server.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	utils.LogFatal(server.Start(":2020").Error())
}