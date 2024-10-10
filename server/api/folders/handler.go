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
	folderName := c.QueryParam("folderName")
	folderPath := c.QueryParam("folderPath")
	folderOwner := c.QueryParam("folderOwner")
	folderID := utils.Hash(folderName)

	// create folder in database
	db := DatabaseGet()
	stmt, err := db.Prepare("INSERT INTO folders (id, name, type, path, owner) VALUES (?, ?, ?, ?, ?)")
	if err != nil {
		utils.LogFatal(err.Error())
	}
	defer stmt.Close()
	_, err = stmt.Exec(folderID, folderName, "personal", folderPath, folderOwner)
	if err != nil {
		utils.LogFatal(err.Error())
	}

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