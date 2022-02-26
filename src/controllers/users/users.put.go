package users

import (
	"encoding/json"
	"log"
	"mm/pkg/src/controllers"
	"mm/pkg/src/db"
	"mm/pkg/src/models"
	"net/http"
)

func put(w http.ResponseWriter, r *http.Request) {
	user := models.Users{}

	json.NewDecoder(r.Body).Decode(&user)

	// stores values from the request
	prev_user := user

	row := db.DB.QueryRow(r.Context(), "select * from users where id = $1", user.ID)

	err := row.Scan(&user.ID, &user.Name)
	if err != nil {
		controllers.ResultNotFound(w, r)
		return
	}

	// if name is not empty
	if prev_user.Name != "" {
		user.Name = prev_user.Name
	}

	_, err = db.DB.Exec(r.Context(), `
		update users set name=$2
		where id = $1
	`, user.ID, user.Name)
	if err != nil {
		log.Println(err)
	}

	json, _ := json.Marshal(models.MsgSingleUser{
		Msg:  "User updated",
		User: user,
	})

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	w.Write(json)
}
