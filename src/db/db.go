package db

import (
	"database/sql"
	"errors"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

// DATABASE the global connection for database
var DATABASE *sql.DB

func init() {
	/*
		Checking if database exists
		Database path should be given at environment variable
		if not given then the default choice is ./database/dev.db
	*/

	DB := os.Getenv("DB")

	if DB == "" {
		log.Println("[DB] Using default database ./database/dev.db")
		DB = "./database/dev.db"
		os.Setenv("DB", DB)
	}

	db, err := sql.Open("sqlite3", DB)

	if err != nil {
		log.Fatalf("[DB] Unable to open sqlite3 DB (%s): %v", DB, err)
	}

	DATABASE = db
}

func StoreImage(uiid, title string, imageData []byte) error {
	_, err := DATABASE.Exec("INSERT INTO images (id, title, image) values (?, ?, ?)", uiid, title, imageData)

	if err != nil {
		return err
	}
	return nil
}

func GetImage(uiid string) ([]byte, error) {
	var image []byte
	err := DATABASE.QueryRow("SELECT image FROM images WHERE id = ?", uiid).Scan(&image)

	if err != nil {
		return nil, errors.New("could not get image from db")
	}

	return image, nil

}
