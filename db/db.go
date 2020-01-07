package db

import (
	"database/sql"
	"fmt"
	"os"
	"os/user"
	"path"
)
import _ "github.com/mattn/go-sqlite3"

var connection *sql.DB
func GetConnection() (conn *sql.DB, err error) {
	currentUser, err := user.Current()
	if err != nil {
		err = fmt.Errorf("unable to open connection: %w", err)
		return
	}
	dbFilePath := path.Join(currentUser.HomeDir, ".local/share/google-photos-sync")
	if _, err = os.Stat(dbFilePath); os.IsNotExist(err) {
		err = os.MkdirAll(dbFilePath, 0755)
		if err != nil {
			err = fmt.Errorf("cannot create db folder %s : %w", dbFilePath, err)
			return
		}
	}
	dbFileName := "photos.db"
	dbFileFullPath := path.Join(dbFilePath, dbFileName)
	conn, err = sql.Open("sqlite3", dbFileFullPath)
	if err != nil {
		err = fmt.Errorf("cannot create db file %s : %w", dbFileFullPath, err)
		return
	}
	connection = conn
	return
}

func initDb() error {
	statement, err := connection.Prepare("CREATE TABLE IF NOT EXISTS tokens (id INTEGER PRIMARY KEY, firstname TEXT, lastname TEXT)")

}
