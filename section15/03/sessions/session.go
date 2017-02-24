package session

import (
	"fmt"
	"net/http"
	"time"

	"github.com/ramirezra/golang-web/section15/03/models"
	"github.com/satori/go.uuid"
)

// Length exported
const Length int = 30

// Users exported
var Users = map[string]models.User{} // user ID, user

// Sessions exported
var Sessions = map[string]models.Session{} // session ID, session

// LastCleaned exported
var LastCleaned time.Time = time.Now()

// GetUser exported
func GetUser(w http.ResponseWriter, req *http.Request) models.User {
	// get cookie
	c, err := req.Cookie("session")
	if err != nil {
		sID := uuid.NewV4()
		c = &http.Cookie{
			Name:  "session",
			Value: sID.String(),
		}

	}
	c.MaxAge = Length
	http.SetCookie(w, c)

	// if the user exists already, get user
	var u models.User
	if s, ok := Sessions[c.Value]; ok {
		s.LastActivity = time.Now()
		Sessions[c.Value] = s
		u = Users[s.UserName]
	}
	return u
}

// AlreadyLoggedIn exported
func AlreadyLoggedIn(w http.ResponseWriter, req *http.Request) bool {
	c, err := req.Cookie("session")
	if err != nil {
		return false
	}
	s, ok := Sessions[c.Value]
	if ok {
		s.LastActivity = time.Now()
		Sessions[c.Value] = s
	}
	_, ok = Users[s.UserName]
	// refresh session
	c.MaxAge = Length
	http.SetCookie(w, c)
	return ok
}

// Clean exported
func Clean() {
	fmt.Println("BEFORE CLEAN") // for demonstration purposes
	Show()                      // for demonstration purposes
	for k, v := range Sessions {
		if time.Now().Sub(v.LastActivity) > (time.Second * 30) {
			delete(Sessions, k)
		}
	}
	LastCleaned = time.Now()
	fmt.Println("AFTER CLEAN") // for demonstration purposes
	Show()                     // for demonstration purposes
}

// Show exported
func Show() {
	fmt.Println("********")
	for k, v := range Sessions {
		fmt.Println(k, v.UserName)
	}
	fmt.Println("")
}
