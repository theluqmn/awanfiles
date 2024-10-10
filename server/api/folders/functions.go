// functions related to folders
package folders

import (
	"main/utils"
)

type Folder struct {
	ID        string
	Name      string
	Type      string
	Path      string
	Owner     string
}

func CreateFolder(folderID, folderName, folderPath, folderOwner string) {
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

func GetFolder(folderID string) (*Folder) {
	db := DatabaseGet()
	stmt, err := db.Prepare("SELECT * FROM folders WHERE id = ?")
	if err != nil {
		utils.LogError(err.Error())
	}
	defer stmt.Close()

	var folder Folder
	err = stmt.QueryRow(folderID).Scan(&folder.ID, &folder.Name, &folder.Type, &folder.Path, &folder.Owner)
	if err != nil {
		utils.LogError(err.Error())
	}

	return &folder
}