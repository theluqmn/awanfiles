package main

import (
	"main/utils"
	"main/server/api/files"

	"net/http"
	"github.com/labstack/echo"
)

func main() {
	// clear terminal
	utils.ClearTerminal()

	// database
	files.DatabaseOpen()

	// server
	server := echo.New()
	server.GET("/", func(c echo.Context) error {
		utils.Log("Hello, World!")
		return c.String(http.StatusOK, "Hello, World!")
	})

	server.POST("/api/files", files.UploadFile)
	server.GET("/api/files/:id", files.DownloadFile)

	// logging
	utils.LogFatal(server.Start(":2020").Error())
}