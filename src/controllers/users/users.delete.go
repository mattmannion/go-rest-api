package users

import (
	"database/sql"
	"encoding/json"
	"log"
	"mm/pkg/src/controllers"
	"mm/pkg/src/db"
	"mm/pkg/src/models"
	"net/http"
)

func delete(w http.ResponseWriter, r *http.Request) {
	user := models.Users{}

	json.NewDecoder(r.Body).Decode(&user)

	// stores values from the request
	post_user := user

	row := db.DB.QueryRow(r.Context(), "select * from users where id = $1", user.ID)
	err := row.Scan(&user.ID, &user.Name)
	if err == sql.ErrNoRows {
		controllers.ResultNotFound(w, r)
		return
	}
	if err != nil {
		log.Println(err)
		return
	}

	// overwrites the select all values
	// if name is not empty
	if post_user.Name != "" {
		user.Name = post_user.Name
	}
	json, _ := json.Marshal(user)

	_, err = db.DB.Exec(r.Context(), `
		update users set name=$2
		where id = $1
	`, user.ID, user.Name)
	if err != nil {
		log.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	w.Write(json)
}
