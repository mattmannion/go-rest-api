package env

import "fmt"

/* Port */
const Port int = 7890

var PortStr = fmt.Sprintf(":%v", Port)

/* Postgres */
const Postgres string = "postgres"
const Pg_dsn string = "postgres://postgres:postgres@localhost/postgres?sslmode=disable"
