package main

import (
	"main/server/api/files"
	"main/server/api/folders"
	"main/utils"

	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	// clear terminal
	utils.ClearTerminal()

	// database
	files.DatabaseOpen()
	folders.DatabaseOpen()

	// server
	server := echo.New()
	// middlewares
	server.Pre(middleware.HTTPSRedirect())
	server.Pre(middleware.AddTrailingSlash())
	server.Use(middleware.Secure())
	server.Use(middleware.CSRF())

	server.GET("/", func(c echo.Context) error {
		utils.Log("Hello, World!")
		return c.String(http.StatusOK, "Hello, World!")
	})

	server.POST("/api/folders", folders.Create)
	server.POST("/api/files", files.UploadFile)
	server.GET("/api/files/:id", files.DownloadFile)

	// logging
	utils.LogFatal(server.Start(":2020").Error())
}