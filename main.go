package main

import (
	"net/http"
	"github.com/labstack/echo"
	"main/server/api/user"
	"main/utils"
)

func main() {
	// Initialize database connection
	db, err := user.DatabaseOpen()
	if err != nil {
		utils.LogFatal("Failed to open database: " + err.Error())
	}
	defer db.Close()

	server := echo.New()
	server.GET("/", func(c echo.Context) error {
		utils.Log("Hello, World!")
		return c.String(http.StatusOK, "Hello, World!")
	})

	server.GET("/api/account/create", user.CreateAccount)

	utils.LogFatal(server.Start(":2020").Error())
}