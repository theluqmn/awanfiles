// functions related to folders
package files

import (
	"fmt"
	"io"
	"os"

	"main/utils"
	"mime/multipart"
)

func UploadFile (file *multipart.FileHeader) {
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()
	
	// variables
	fileSize := (file.Size / 1024) / 1024 // unit: MB
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

	// record file in database
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
	
    if _, err = io.Copy(dst, src); err != nil {
		utils.LogError(err.Error())
    }
}