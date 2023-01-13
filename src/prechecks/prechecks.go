package prechecks

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

/*
Pre Checks are the checks which will check for the existence of
some critical parts of program like db
*/

func init() {
	log.Println("[PRE CHECK] Running Pre Checks...")

	/*
		Checking if database exists
		Database path should be given at environment variable
		if not given then the default choice is ./database/dev.db
	*/
	DB := os.Getenv("DB")
	if DB == "" {
		log.Println("[PRE CHECK] Using default Database (./database/dev.db) ")
		DB = "./database/dev.db"
	}

	log.Println("[PRE CHECK] Checking if Database exists in ./database...")
	if _, err := os.Stat(DB); err != nil {
		log.Fatal("[PRE CHECK] Database does not exist, make sure you have a db file in ./database")
	} else {
		log.Println("[PRE CHECK] Database exist")
	}

	// Check if the table has required table
	log.Println("[PRE CHECK] Checking if Database Table (images) exists...")

	db, err := sql.Open("sqlite3", DB)

	if err != nil {
		log.Fatalf("[PRE CHECK] %v", err.Error())
	}

	table, err_table := db.Query("SELECT name FROM sqlite_master WHERE type='table' AND name='images'")

	if err_table != nil {
		log.Fatalf("[PRE CHECK] %v ", err_table.Error())
	}

	if table.Next() {
		log.Println("[PRE CHECK] Images table exists")
	} else {
		log.Fatal("[PRE CHECK] Images table does not exist")
	}

	// check if the table has specific columns
	log.Println("[PRE CHECK] Checking if Images table has all the required columns...")
	checkColumnExist("id", db)
	checkColumnExist("title", db)
	checkColumnExist("blob", db)
	log.Println("[PRE CHECK] Images table has all the required columns")

	log.Println("[PRE CHECK] Passed all checks")

}

func checkColumnExist(column_name string, db *sql.DB) {
	columns, err_columns := db.Query(fmt.Sprintf("SELECT COUNT(*) AS CNTREC FROM pragma_table_info('images') WHERE name='%s'", column_name))

	if err_columns != nil {
		log.Fatalf("[PRE CHECK] %v ", err_columns.Error())
	}

	if columns.Next() {
		var count int
		columns.Scan(&count)
		if count == 0 {
			log.Fatalf("[PRE CHECK] Images table does not have %s column", column_name)
		}
	} else {
		log.Fatalf("[PRE CHECK] Images table does not have %s column", column_name)
	}
}
