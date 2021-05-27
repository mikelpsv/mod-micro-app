package app

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
)

var Db *sql.DB

func InitDb(host, dbname, dbuser, dbpass string) {
	var err error

	connStr := fmt.Sprintf("host=%s dbname=%s user=%s password=%s sslmode=disable", host, dbname, dbuser, dbpass)
	Db, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	err = Db.Ping()
	if err != nil {
		log.Fatal(err)
	}

}

func Close() {
	Db.Close()
}

func Install(loadTestData bool) {

}
