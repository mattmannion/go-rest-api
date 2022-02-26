package users

import (
	"encoding/json"
	"mm/pkg/src/controllers"
	"mm/pkg/src/db"
	"mm/pkg/src/models"
	"net/http"
)

func get_one(w http.ResponseWriter, r *http.Request, id int) {
	row := db.DB.QueryRow(r.Context(), "select * from users where id = $1", id)

	user := models.Users{}
	err := row.Scan(&user.ID, &user.Name)
	if err != nil {
		controllers.ResultNotFound(w, r)
		return
	}

	json, _ := json.Marshal(models.MsgSingleUser{
		Msg:  "User retrieved",
		User: user,
	})

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	w.Write(json)
}
