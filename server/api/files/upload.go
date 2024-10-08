package files

import (
	"io"
	"os"
	"main/utils"

	"net/http"
	"github.com/labstack/echo"
)

func UploadFile(c echo.Context) error {
    // Read form file
    file, err := c.FormFile("file")
    if err != nil {
        return err
    }

    // File source
    src, err := file.Open()
    if err != nil {
        return err
    }
    defer src.Close()

    // Destination
    dst, err := os.Create("files/" + file.Filename)
    if err != nil {
        return err
    }
    defer dst.Close()

    // Copy at server
    if _, err = io.Copy(dst, src); err != nil {
        return err
    }

	utils.Log("New file uploaded")
    return c.String(http.StatusOK, "File uploaded successfully")
}