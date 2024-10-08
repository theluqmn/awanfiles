// function for handling file download to clients
package files

import (
    "os"
	"fmt"
	
	"main/utils"

    "net/http"
    "github.com/labstack/echo"
)

func DownloadFile(c echo.Context) error {
    // Get the file ID from the request parameters
    file := c.Param("id")
	fileID := utils.GetFileName(file)
	fileFormat := utils.GetFileFormat(file)

	path := "files/" + fileID + fileFormat
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return c.String(http.StatusNotFound, "File not found")
	}

	utils.Log(fmt.Sprintf("File '%s' downloaded", fileID))
	return c.Attachment(path, file)
}