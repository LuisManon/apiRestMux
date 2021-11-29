package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/LuisManon/apiRestMux/pkg/mocks"
	"github.com/gorilla/mux"
)

func GetUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	for _, user := range mocks.User {
		if user.Id == id {
			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(user)
			break
		}
	}
}
