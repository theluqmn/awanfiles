// create folders
package folders

import (
	"fmt"

	"main/utils"
	
	"net/http"
	"github.com/labstack/echo"
)

func CreateFolder(c echo.Context) error {
	// variables
	folderName := c.QueryParam("folderName")
	folderPath := c.QueryParam("folderPath")
	folderOwner := c.QueryParam("folderOwner")
	folderID := utils.Hash(folderName)

	RecordFolder(folderID, folderName, folderPath, folderOwner)

	return c.String(http.StatusOK, fmt.Sprintf("Created new folder '%s' with ID: %s", folderName, folderID))
}

// record folder in database
func RecordFolder(folderID string, folderName string, folderPath string, folderOwner string) {
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
}