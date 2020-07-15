package factory

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

func GetConnection() *sql.DB {
	dbHost := os.Getenv("DB_HOST")
	dbUser := os.Getenv("DB_USER")
	bdPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	dbinfo := fmt.Sprintf("host=%s user=%s password=%s dbname=%s",
		dbHost, dbUser, bdPassword, dbName)

	db, err := sql.Open("postgres", dbinfo)
	checkErr(err)
	return db
}

func GetConnectionSchema(schema string) *sql.DB {
	dbHost := os.Getenv("DB_HOST")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	dbinfo := fmt.Sprintf("host=%s user=%s password=%s dbname=%s search_path=%s,public",
		dbHost, dbUser, dbPassword, dbName, schema)

	db, err := sql.Open("postgres", dbinfo)
	checkErr(err)
	return db
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
