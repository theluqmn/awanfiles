// functions for handling file uploads.
package files

import (
	"fmt"
	"io"
	"os"

	"main/utils"
	
	"net/http"
	"github.com/labstack/echo"
)

func UploadFile(c echo.Context) error {
	// read from file
    file, err := c.FormFile("file")
    if err != nil {
        return err
    }

	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	// variables
	fileSize := (file.Size / 1024) / 1024 // MB
	fileName := utils.GetFileName(file.Filename)
	fileFormat := utils.GetFileFormat(file.Filename)
	fmt.Printf("Format: %s\n", fileFormat)
	fileID := utils.Hash(fileName)

	// write to server
    dst, err := os.Create("files/" + fileID + fileFormat)
    if err != nil {
        utils.LogError(err.Error())
    }
    defer dst.Close()
	
    if _, err = io.Copy(dst, src); err != nil {
		utils.LogError(err.Error())
    }

	// record file
	recordFile(fileID, fileName, fileFormat, 1)

	utils.Log(fmt.Sprintf("New file '%s' (format: %s) uploaded with size: %d MB\nas: %s", fileName, fileFormat, fileSize, fileID))
    return c.String(http.StatusOK, fmt.Sprintf("Uploaded successfully as %s", fileID))
}

// record file in database
func recordFile(id string, name string, format string, owner int) {
	db := DatabaseGet()
	stmt, err := db.Prepare("INSERT INTO files (id, name, format, owner) VALUES (?, ?, ?, ?)")
	if err != nil {
		utils.LogFatal(err.Error())
	}
	defer stmt.Close()
	_, err = stmt.Exec(id, name, format, owner)
	if err != nil {
		utils.LogFatal(err.Error())
	}

	utils.Log(fmt.Sprintf("Recorded file '%s' in database", id))
}