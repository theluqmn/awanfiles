package main

import (
	"net/http"
	"github.com/labstack/echo"
)

func main() {
	server := echo.New()
	server.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	server.Logger.Fatal(server.Start(":2020"))
}