package users

import (
	"mm/pkg/src/controllers"
	"net/http"
)

func get(w http.ResponseWriter, r *http.Request) {
	controllers.GetByID(get_one, get_all, w, r)
}
