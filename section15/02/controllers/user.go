package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/ramirezra/golang-web/section15/02/models"
	uuid "github.com/satori/go.uuid"
)

// GetUser exported
func (c Controller) GetUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	// Grab id
	id := p.ByName("id")

	// Fetch user
	u := uc.session[id]

	// Marshal provided interface into JSON structure
	uj, err := json.Marshal(u)
	if err != nil {
		fmt.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK) // 200
	fmt.Fprintf(w, "%s\n", uj)
}

// CreateUser exported
func (uc UserController) CreateUser(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	u := models.User{}

	json.NewDecoder(r.Body).Decode(&u)

	// create ID
	u.Id = uuid.NewV4().String()

	// store the user
	uc.session[u.Id] = u
	models.StoreUsers(uc.session)
	uj, err := json.Marshal(u)
	if err != nil {
		fmt.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated) // 201
	fmt.Fprintf(w, "%s\n", uj)
}

// DeleteUser exported
func (uc UserController) DeleteUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id := p.ByName("id")

	// Delete user
	delete(uc.session, id)
	models.StoreUsers(uc.session)
	w.WriteHeader(http.StatusOK) // 200
	fmt.Fprint(w, "Deleted user", id, "\n")
}
