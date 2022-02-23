package db

import (
	"database/sql"
	"fmt"
	"log"
	"mm/pkg/src/env"
)

var DB *sql.DB
var err error

func Init() {
	DB, err = sql.Open(env.Postgres, env.Pg_dsn)

	if err != nil {
		log.Fatal(err)
	}
	if err = DB.Ping(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("connected to db")
}
