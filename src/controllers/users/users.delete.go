package users

import (
	"encoding/json"
	"mm/pkg/src/controllers"
	"mm/pkg/src/db"
	"mm/pkg/src/models"
	"net/http"
)

func delete(w http.ResponseWriter, r *http.Request) {
	user := models.Users{}

	json.NewDecoder(r.Body).Decode(&user)

	row := db.DB.QueryRow(r.Context(), "select * from users where id = $1", user.ID)

	err := row.Scan(&user.ID, &user.Name)
	if err != nil {
		controllers.ResultNotFound(w, r)
		return
	}

	db.DB.Exec(r.Context(), "delete from users where id=$1", user.ID)

	json, _ := json.Marshal(models.MsgSingleUser{
		Msg:  "User deleted",
		User: user,
	})

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	w.Write(json)
}
