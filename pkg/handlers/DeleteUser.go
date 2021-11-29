package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/LuisManon/apiRestMux/pkg/mocks"
	"github.com/gorilla/mux"
)

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	for index, user := range mocks.User {
		if user.Id == id {
			mocks.User = append(mocks.User[:index], mocks.User[index+1:]...)

			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode("Deleted")
			break
		}
	}
}