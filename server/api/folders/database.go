// functions for handling database operations related to folders.
package folders

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
		db, err = sql.Open("sqlite3", "server/database/folders.db")
		if err != nil {
			utils.LogError(err.Error())
		}
		err = db.Ping() // Ensure the database connection is established
		if err != nil {
			utils.LogError(err.Error())
		}
		_, err = db.Exec("CREATE TABLE IF NOT EXISTS folders (id TEXT PRIMARY KEY, name TEXT, type TEXT, path TEXT, owner INTEGER)")
	})
	return db, err
}

func DatabaseGet() *sql.DB {
	return db
}