// functions for handling database operations related to files.
package files

import (
	"database/sql"
	"main/utils"

	_ "github.com/mattn/go-sqlite3"

	"sync"
)

var (
	db *sql.DB
	once sync.Once
)

func DatabaseOpen() (*sql.DB, error) {
	var err error
	once.Do(func() {
		db, err = sql.Open("sqlite3", "server/database/files.db")
		if err != nil {
			utils.LogError(err.Error())
		}
		err = db.Ping() // Ensure the database connection is established
		if err != nil {
			utils.LogError(err.Error())
		}
		_, err = db.Exec("CREATE TABLE IF NOT EXISTS files (id TEXT PRIMARY KEY, folder_id TEXT, size FLOAT)")
	})
	return db, err
}

func DatabaseGet() *sql.DB {
	return db
}