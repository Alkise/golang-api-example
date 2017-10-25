package v1

import (
	"encoding/json"
	"net/http"

	"models"

	"github.com/gorilla/mux"
)

// UsersHandler Users page handler
func UsersHandler(w http.ResponseWriter, r *http.Request) {

	// newUser := models.User{Email: "looi@looi.co", FirstName: "User", LastName: "Second"}
	// _, err := newUser.Save()
	// if err != nil {
	// 	http.Error(w, err.Error(), http.StatusInternalServerError)
	// }

	// users := &models.UserCollection{}
	users, err := models.AllUsers()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	jsonUsers, err := json.Marshal(*users)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonUsers)
}

// UserHandler Single users page handler
func UserHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	user, err := models.FindUser(vars["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	jsonUser, err := json.Marshal(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonUser)
}
