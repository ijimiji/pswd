package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"pswd/pkg/config"

	_ "github.com/mattn/go-sqlite3" // Import go-sqlite3 library
)

var cfg = config.Get()

// Create home directory for database, configs and etc
func ensureHomeFolder() {
	if _, err := os.Stat(cfg.HomeFolder); os.IsNotExist(err) {
		os.Mkdir(cfg.HomeFolder, 0700)
	}
}

// Create default database if it doesn't exists
// and create password table
func ensureDatabase() *sql.DB {
	dbFilename := fmt.Sprintf("%s/%s", cfg.HomeFolder, "pswd.db")
	if _, err := os.Stat(dbFilename); os.IsNotExist(err) {
		file, err := os.Create(dbFilename)
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()
	}
	database, _ := sql.Open("sqlite3", dbFilename)
	database.Exec(`CREATE TABLE IF NOT EXISTS PASSWORDS (KEY TEXT, PASSWORD TEXT);`)
	return database
}

func GetDatabase() *sql.DB {
	ensureHomeFolder()
	return ensureDatabase()
}
