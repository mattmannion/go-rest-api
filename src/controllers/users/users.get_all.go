package users

import (
	"encoding/json"
	"log"
	"mm/pkg/src/db"
	"mm/pkg/src/models"
	"net/http"
)

func get_all(w http.ResponseWriter, r *http.Request) {
	rows, err := db.DB.Query(r.Context(), "select * from users order by id")
	if err != nil {
		log.Println(err)
	}
	defer rows.Close()

	users := make([]models.Users, 0)

	for rows.Next() {
		user := models.Users{}
		err := rows.Scan(&user.ID, &user.Name)
		if err != nil {
			log.Println(err)
			return
		}

		users = append(users, user)
	}

	json, _ := json.Marshal(models.MsgManyUsers{
		Msg:   "Users retrieved",
		Users: users,
	})

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	w.Write(json)
}
