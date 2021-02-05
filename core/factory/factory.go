package factory

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
	"github.com/spf13/viper"
)

func GetConnection() *sql.DB {
	dbHost := viper.GetString(`database.host`)
	dbUser := viper.GetString(`database.user`)
	dbPassword := viper.GetString(`database.pass`)
	dbName := viper.GetString(`database.name`)

	dbinfo := fmt.Sprintf("host=%s user=%s password=%s dbname=%s",
		dbHost, dbUser, dbPassword, dbName)

	db, err := sql.Open("postgres", dbinfo)
	checkErr(err)

	return db
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
