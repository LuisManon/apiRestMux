package handlers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/LuisManon/apiRestMux/pkg/mocks"
	"github.com/LuisManon/apiRestMux/pkg/models"
	"github.com/gorilla/mux"
)

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Fatalln(err)
	}

	var updatedUser models.User
	json.Unmarshal(body, &updatedUser)

	for index, user := range mocks.User {
		if user.Id == id {
			user.Name = updatedUser.Name
			user.Email = updatedUser.Email
			user.Password = updatedUser.Password

			mocks.User[index] = user

			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode("Updated")
			break
		}
	}
}
