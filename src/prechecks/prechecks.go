package prechecks

import (
	"log"
	"os"
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

	log.Println("[PRE CHECK] Checking if Database exists in ./database")
	if _, err := os.Stat(DB); err != nil {
		log.Fatal("[PRE CHECK] Database does not exist, make sure you have a db file in ./database")
	} else {
		log.Println("[PRE CHECK] Database exist")
	}
}
