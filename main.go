package main

import (
	"main/utils"

	"net/http"
	"github.com/labstack/echo"
)

func main() {
	server := echo.New()
	server.GET("/", func(c echo.Context) error {
		utils.Log("Hello, World!")
		return c.String(http.StatusOK, "Hello, World!")
	})

	server.GET("/api/files/:user", )

	utils.LogFatal(server.Start(":2020").Error())
}