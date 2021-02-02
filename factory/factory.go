package factory

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

type Connection struct {
	dbinfo string
}

func GetConnection() Connection {
	connection := Connection{}

	dbHost := os.Getenv("DB_HOST")
	dbUser := os.Getenv("DB_USER")
	bdPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	connection.dbinfo = fmt.Sprintf("host=%s user=%s password=%s dbname=%s",
		dbHost, dbUser, bdPassword, dbName)
	return connection

	// sql.Open("postgres", dbinfo)
}

func (connection Connection) Open() *sql.DB {
	db, err := sql.Open("postgres", connection.dbinfo)

	checkErr(err)

	return db
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
