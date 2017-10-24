package v1

import (
	"encoding/json"
	"net/http"

	"models"

	"github.com/gorilla/mux"
)

// UsersHandler Users page handler
func UsersHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// newUser := models.User{Email: "looi@looi.co", FirstName: "User", LastName: "Second"}
	// _, err := newUser.Save()
	// panicOnErr(err)

	// users := &models.UserCollection{}
	users, err := models.AllUsers()
	panicOnErr(err)

	jsonUsers, err := json.Marshal(users)
	panicOnErr(err)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonUsers)
}

// UserHandler Single users page handler
func UserHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	user, err := models.FindUser(vars["id"])
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	jsonUser, err := json.Marshal(user)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(jsonUser)
}
