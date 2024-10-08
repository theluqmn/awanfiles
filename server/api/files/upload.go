package files

import (
	"fmt"
	"io"
	"os"
	"time"

	"main/utils"

	"net/http"
	"github.com/labstack/echo"
)

func UploadFile(c echo.Context) error {
    file, err := c.FormFile("file") // read form file
    if err != nil {
        return err
    }

	fileSize := (file.Size / 1024) / 1024 // MB
	fileName := file.Filename
	hashName := createFileRecord(fileName, 1)

    src, err := file.Open() // file source
    if err != nil {
        return err
    }
    defer src.Close()

    dst, err := os.Create("files/" + hashName) // destination
    if err != nil {
        return err
    }
    defer dst.Close()
    if _, err = io.Copy(dst, src); err != nil {
        return err
    }

	utils.Log(fmt.Sprintf("New file '%s' uploaded with size: %d MB\n", fileName, fileSize))
    return c.String(http.StatusOK, fmt.Sprintf("Uploaded successfully as %s", hashName))
}

func createFileRecord(name string, owner int) (hashname string) {
	// Create a hash of the file name as unique identifier
	hashname = utils.Hash(name)

	db := DatabaseGet()
	stmt, err := db.Prepare("INSERT INTO files (id, owner) VALUES (?,?)")
	if err != nil {
		utils.LogFatal("Failed to prepare statement")
	}
	defer stmt.Close()

	currentTime := time.Now()
	_, err = stmt.Exec(hashname, owner)
	if err != nil {
		utils.LogFatal("Failed to create file")
	}

	return hashname
}