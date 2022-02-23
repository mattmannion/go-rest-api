package db

import (
	"database/sql"
	"fmt"
	"log"
	"mm/pkg/src/env"
)

var PG *sql.DB
var err error

func Init() {
	PG, err = sql.Open(env.Postgres, env.Pg_dsn)

	if err != nil {
		log.Fatal(err)
	}
	if err = PG.Ping(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("connected to db")
}
