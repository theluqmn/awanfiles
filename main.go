package main

import (
	"main/utils"
	"main/server/api/user"
	"main/server/api/files"

	"net/http"
	"github.com/labstack/echo"
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

	// server.GET("/api/account/create", user.CreateAccount)
	server.POST("/api/file/upload", files.UploadFile)

	// Logging
	utils.LogFatal(server.Start(":2020").Error())
}