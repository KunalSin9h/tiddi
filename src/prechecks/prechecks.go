package prechecks

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"github.com/znip-in/tiddi/src/db"
)

/*
Pre Checks are the checks which will check for the existence of
some critical parts of program like db
*/

func RUN_PRECHECKS() {
	log.Println("[PRE CHECK] Running Pre Checks...")

	log.Println("[PRE CHECK] Checking if Database exists in ./database...")
	if _, err := os.Stat(os.Getenv("DB")); err != nil {
		log.Fatal("[PRE CHECK] Database does not exist, make sure you have a db file in ./database")
	} else {
		log.Println("[PRE CHECK] Database exist")
	}

	// Check if the db has required table
	log.Println("[PRE CHECK] Checking if Database Table (images) exists...")

	table, err_table := db.DATABASE.Query("SELECT name FROM sqlite_master WHERE type='table' AND name='images'")

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
	checkColumnExist("id", db.DATABASE)
	checkColumnExist("title", db.DATABASE)
	checkColumnExist("image", db.DATABASE)
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
