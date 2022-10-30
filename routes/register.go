package routes

import (
	"encoding/json"
	"net/http"

	"github.com/rodrigoSolano/twitter_go/db"
	"github.com/rodrigoSolano/twitter_go/models"
)

// Register is a function to register a new user
func Register(w http.ResponseWriter, r *http.Request) {
	var t models.User
	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w, "Incorrect data "+err.Error(), 400)
		return
	}
	if len(t.Email) == 0 {
		http.Error(w, "Email is required", 400)
		return
	}
	if len(t.Password) < 6 {
		http.Error(w, "Password must have at least 6 characters", 400)
		return
	}
	_, found, _ := db.CheckUserExist(t.Email)
	if found == true {
		http.Error(w, "User already exist", 400)
		return
	}
	_, status, err := db.InsertUser(t)
	if err != nil {
		http.Error(w, "Error while trying to register user "+err.Error(), 400)
		return
	}
	if status == false {
		http.Error(w, "Could not register user", 400)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
