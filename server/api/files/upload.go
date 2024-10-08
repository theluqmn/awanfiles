package files

import (
	"main/utils"
	"net/http"
	"github.com/labstack/echo"
)

func UploadFile(c echo.Context) error {
	db := DatabaseGet()
	
}