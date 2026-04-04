package database

import (
	"database/sql"
	"gpm/module/util"
	"path/filepath"

	_ "modernc.org/sqlite"
)

var initQueries []string = []string{
	"CREATE TABLE IF NOT EXISTS logfile (name TEXT, filename TEXT);",
	`CREATE TABLE IF NOT EXISTS "logfile-main" (name TEXT, filename TEXT);`,
}

func OpenDB() (*sql.DB, error) {
	dbPath, err := getDBPath()
	if err != nil {
		return nil, err
	}

	db, err := sql.Open("sqlite", dbPath)
	if err != nil {
		return nil, err
	}

	for _, query := range initQueries {
		_, err = db.Exec(query)
		if err != nil {
			return nil, err
		}
	}

	return db, nil
}

func getDBPath() (string, error) {
	homeDir, err := util.GetHomeDirPath()
	if err != nil {
		return "", err
	}
	return filepath.Join(homeDir, "main.db"), nil
}
