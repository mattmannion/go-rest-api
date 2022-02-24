package db

import (
	"context"
	"fmt"
	"log"
	"mm/pkg/src/env"

	"github.com/jackc/pgx/v4/pgxpool"
)

var DB *pgxpool.Pool
var err error

func Init() {
	DB, err = pgxpool.Connect(context.Background(), env.Pg_dsn)
	// DB, err = sql.Open(env.Postgres, env.Pg_dsn)

	if err != nil {
		log.Fatal(err)
	}
	if err = DB.Ping(context.Background()); err != nil {
		log.Fatal(err)
	}

	fmt.Println("connected to db")
}
