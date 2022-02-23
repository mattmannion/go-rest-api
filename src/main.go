package main

import (
	"fmt"
	"mm/pkg/src/db"
	"mm/pkg/src/env"
	"mm/pkg/src/router"
	"net/http"

	_ "github.com/lib/pq"
)

func main() {
	db.Init()
	fmt.Printf("live @ http://localhost%s\n", env.PortStr)
	http.ListenAndServe(env.PortStr, router.Router)
}
