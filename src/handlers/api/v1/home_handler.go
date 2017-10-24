package v1

import (
	"net/http"
)

// APIVersion Version of API
const APIVersion string = "v1"

// HomeHandler Home page handler
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("API " + APIVersion + " Home page"))
}
