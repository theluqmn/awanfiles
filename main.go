package main

import (
	"net/http"
	"github.com/labstack/echo"
	"main/server/api/user"
	"main/utils"
)

func main() {
	// Database
	db, err := user.DatabaseOpen()
	if err != nil {
		utils.LogFatal("Failed to open database: " + err.Error())
	}
	defer db.Close()

	// Server
	server := echo.New()
	server.GET("/", func(c echo.Context) error {
		utils.Log("Hello, World!")
		return c.String(http.StatusOK, "Hello, World!")
	})

	server.GET("/api/account/create", user.CreateAccount)

	// Logging
	utils.LogFatal(server.Start(":2020").Error())
}