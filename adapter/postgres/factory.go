package postgres

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v4/pgxpool"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
)

func GetConnection(context context.Context) *pgxpool.Pool {
	databaseUrl := viper.GetString("database.url")
	conn, err := pgxpool.Connect(context, databaseUrl)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}

	return conn
}
