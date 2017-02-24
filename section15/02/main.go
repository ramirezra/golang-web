package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/ramirezra/golang-web/section15/02/controllers"
	"github.com/ramirezra/golang-web/section15/02/models"
)

func main() {
	r := httprouter.New()
	// Get a UserController instance
	uc := controllers.NewUserController(getSession())
	r.GET("/user/:id", uc.GetUser)
	r.POST("/user", uc.CreateUser)
	r.DELETE("/user/:id", uc.DeleteUser)
	http.ListenAndServe("localhost:8080", r)
}

func getSession() map[string]models.User {
	return models.LoadUsers()
}
