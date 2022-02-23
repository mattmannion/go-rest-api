package users

import (
	"encoding/json"
	"log"
	"mm/pkg/src/db"
	"mm/pkg/src/models"
	"net/http"
)

func get_all(w http.ResponseWriter, r *http.Request) {
	rows, err := db.DB.Query("select * from users")
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

	json, err := json.Marshal(users)

	if err != nil {
		log.Println(err)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(500)
		w.Write([]byte("error-500"))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	w.Write(json)
}
