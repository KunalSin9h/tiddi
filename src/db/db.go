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

func SETUP_DB() {
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

func GetImageBytes(uiid string) ([]byte, error) {
	var image []byte
	err := DATABASE.QueryRow("SELECT image FROM images WHERE id = ?", uiid).Scan(&image)

	if err != nil {
		return nil, errors.New("could not get image from db")
	}

	return image, nil

}

func GetImageDetails(uiid string) (string, []byte, error) {
	var title string
	var image []byte
	err := DATABASE.QueryRow("SELECT title, image FROM images WHERE id = ?", uiid).Scan(&title, &image)

	if err != nil {
		return "", nil, errors.New("could not get image details from db")
	}

	return title, image, nil

}

func DeleteImageFormDB(uiid string) error {

	_, err := DATABASE.Exec("DELETE FROM images where id=?", uiid)

	/*
		err will not caused due to absence of that image
			sql will not delete any think if it does't exist

		So error here means something got wrong with db, itself
	*/

	if err != nil {
		return errors.New("db error, while deleting entry")
	}

	return nil

}

func UpdateTitle(uiid, title string) error {
	_, err := DATABASE.Exec("UPDATE images SET title=? WHERE id=?", title, uiid)

	if err != nil {
		return errors.New("db error, while updating title")
	}

	return nil
}

func UpdateImageData(uiid string, image []byte) error {
	_, err := DATABASE.Exec("UPDATE images SET image=? WHERE id=?", image, uiid)

	if err != nil {
		return errors.New("db error, while updating image data")
	}

	return nil
}
