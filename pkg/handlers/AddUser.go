package handlers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"

	"github.com/tutorials/go/crud/pkg/mocks"
	"github.com/tutorials/go/crud/pkg/models"
)

func AddUser(w http.ResponseWriter, r *http.Request) {

	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Fatalln(err)
	}

	var user models.User
	json.Unmarshal(body, &user)

	user.Id = rand.Intn(100)
	mocks.User = append(mocks.User, user)

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode("Created")
}
