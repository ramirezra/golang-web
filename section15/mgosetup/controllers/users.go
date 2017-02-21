package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"github.com/julienschmidt/httprouter"
	"github.com/ramirezra/golang-web/section15/mgosetup/models"
)

// UserController exported to main program
type UserController struct {
	session *mgo.Session
}

// NewUserController interface
func NewUserController(s *mgo.Session) *UserController {
	return &UserController{s}
}

// GetUser needs to be capitalized for export back to main.go
func (uc UserController) GetUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	// Define user
	// u := models.User{
	// 	Name:   "James Bond",
	// 	Gender: "male",
	// 	Age:    32,
	// 	Id:     p.ByName("id"),
	// }
	id := p.ByName("id")

	// Verify id is ObjectID
	if !bson.IsObjectIdHex(id) {
		w.WriteHeader(http.StatusNotFound) //404
		return
	}

	// ObjectIdHex returns an ObjectIdHex
	oid := bson.ObjectIdHex(id)

	// define model
	u := models.User{}

	// Fetch users
	if err := uc.session.DB("go-web-dev-db").C("users").FindId(oid).One(&u); err != nil {
		w.WriteHeader(404)
		return
	}

	uj, err := json.Marshal(u)
	if err != nil {
		fmt.Println(err)
	}

	// Write headers
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK) //200
	fmt.Fprintf(w, "%s\n", uj)
}

// CreateUser capitalized for export
func (uc UserController) CreateUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	// define composite literal
	u := models.User{}

	// encode/decode for sending/receiving JSON to/from a stream
	json.NewDecoder(r.Body).Decode(&u)

	// Change Id
	u.Id = bson.NewObjectId()

	// Store user in mongodb
	uc.session.DB("go-web-dev-db").C("users").Insert(u)

	// marshal/unmarshal for having JSON assigned to a variable
	uj, err := json.Marshal(u)
	if err != nil {
		fmt.Println(err)
	}
	// Write content-type, status code, and payload
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated) // 201
	fmt.Fprintf(w, "%s\n", uj)
}

// DeleteUser capitalized for export
func (uc UserController) DeleteUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	// TODO: write code to delete user
	id := p.ByName("id")

	if !bson.IsObjectIdHex(id) {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	oid := bson.ObjectIdHex(id)

	// Delete user
	if err := uc.session.DB("go-web-dev-db").C("users").RemoveId(oid); err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusOK) // 200
	fmt.Fprint(w, "Write code to delete user\n")
}
