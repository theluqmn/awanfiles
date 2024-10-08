package user

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"

	"sync"
)

var (
	db *sql.DB
	once sync.Once
)

func DatabaseOpen(){
	var err error
	once.Do(func() {
		db, err = sql.Open("sqlite3", "server/database/users.db")
		if err != nil {
			return
		}
		err = db.Ping() // Ensure the database connection is established
		if err != nil {
			return
		}
		_, err = db.Exec("CREATE TABLE IF NOT EXISTS users (id INTEGER PRIMARY KEY, username TEXT, password TEXT)")
	})
}

func DatabaseGet() *sql.DB {
	return db
}