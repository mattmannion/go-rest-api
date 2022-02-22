package main

import (
	"fmt"
	"mm/pkg/src/router"
	"net/http"
)

func main() {
	fmt.Printf("live @ http://localhost%s\n", PortStr)
	http.ListenAndServe(PortStr, router.Router)
}
