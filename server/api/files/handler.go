// functions for handling database operations related to files.
package files

import (
	"database/sql"
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
			return
		}
		err = db.Ping() // Ensure the database connection is established
		if err != nil {
			return
		}
		_, err = db.Exec("CREATE TABLE IF NOT EXISTS files (id TEXT PRIMARY KEY, name TEXT, format TEXT, owner INTEGER)")
	})
	return db, err
}

func DatabaseGet() *sql.DB {
	return db
}