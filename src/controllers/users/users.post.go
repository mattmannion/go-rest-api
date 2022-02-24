package users

import (
	"encoding/json"
	"log"
	"mm/pkg/src/db"
	"mm/pkg/src/models"
	"net/http"
)

func post(w http.ResponseWriter, r *http.Request) {
	user := models.Users{}

	json.NewDecoder(r.Body).Decode(&user)

	if user.Name == "" {
		error := struct {
			Error string       `json:"msg"`
			User  models.Users `json:"user"`
		}{
			Error: "Data requires the following format:",
			User: models.Users{
				Name: "Enter a name here",
			},
		}

		json, _ := json.Marshal(error)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(400)
		w.Write(json)
		return
	}

	_, err := db.DB.Exec(r.Context(), `
		insert into users(name)
		values($1)
	`, user.Name)
	if err != nil {
		log.Println(err)
	}

	// -- start - brute force last id retrieval
	// very slow...
	rows, err := db.DB.Query(r.Context(), "select * from users")
	if err != nil {
		log.Println(err)
	}
	defer rows.Close()

	//
	var id int = 0
	for rows.Next() {
		id++
	}
	user.ID = id
	// -- end - brute force last id retrieval

	json, err := json.Marshal(user)
	if err != nil {
		log.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	w.Write(json)
}