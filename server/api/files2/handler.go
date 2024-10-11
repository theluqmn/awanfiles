// API handler for file-related actions
package files

import (
	"fmt"

	"main/utils"
	
	"net/http"
	"github.com/labstack/echo"
)

func Upload(c echo.Context) error {
	// variables
	file, err := c.FormFile("file")
	if err != nil {
		utils.LogError(err.Error())
	}

	UploadFile(file)
}