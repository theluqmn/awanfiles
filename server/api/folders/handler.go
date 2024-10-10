// API handler for folder-related actions
package folders

import (
	"fmt"

	"main/utils"
	
	"net/http"
	"github.com/labstack/echo"
)

func Create(c echo.Context) error {
	// variables
	folderName := c.QueryParam("name")
	folderPath := c.QueryParam("path")
	folderOwner := c.QueryParam("owner")
	folderID := utils.Hash(folderName)

	CreateFolder(folderID, folderName, folderPath, folderOwner)

	utils.Log("New folder created")
	return c.String(http.StatusOK, fmt.Sprintf("Created new folder '%s' with ID: %s", folderName, folderID))
}

func Get(c echo.Context) error {
	// variables
	folderID := c.Param("id")
	folder := GetFolder(folderID)
	
	utils.Log("Folder fetched")
	return c.JSON(http.StatusOK, folder)
}