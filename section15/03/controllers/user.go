package controllers

import (
	"net/http"
	"time"

	"github.com/ramirezra/golang-web/section15/03/models"
	session "github.com/ramirezra/golang-web/section15/03/sessions"
	"github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
)

// SignUp exported
func (c Controller) SignUp(w http.ResponseWriter, req *http.Request) {
	if session.AlreadyLoggedIn(w, req) {
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}
	var u models.User
	// process form submission
	if req.Method == http.MethodPost {
		// get form values
		un := req.FormValue("username")
		p := req.FormValue("password")
		f := req.FormValue("firstname")
		l := req.FormValue("lastname")
		r := req.FormValue("role")
		// username taken?
		if _, ok := session.Users[un]; ok {
			http.Error(w, "Username already taken", http.StatusForbidden)
			return
		}
		// create session
		sID := uuid.NewV4()
		c := &http.Cookie{
			Name:  "session",
			Value: sID.String(),
		}
		c.MaxAge = session.Length
		http.SetCookie(w, c)
		session.Sessions[c.Value] = models.Session{un, time.Now()}
		// store user in dbUsers
		bs, err := bcrypt.GenerateFromPassword([]byte(p), bcrypt.MinCost)
		if err != nil {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}
		u = models.User{un, bs, f, l, r}
		session.Users[un] = u
		// redirect
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}
	session.Show() // for demonstration purposes
	c.tpl.ExecuteTemplate(w, "signup.gohtml", u)
}

// Login exported
func (c Controller) Login(w http.ResponseWriter, req *http.Request) {
	if session.AlreadyLoggedIn(w, req) {
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}
	var u models.User
	// process form submission
	if req.Method == http.MethodPost {
		un := req.FormValue("username")
		p := req.FormValue("password")
		// is there a username?
		u, ok := session.Users[un]
		if !ok {
			http.Error(w, "Username and/or password do not match", http.StatusForbidden)
			return
		}
		// does the entered password match the stored password?
		err := bcrypt.CompareHashAndPassword(u.Password, []byte(p))
		if err != nil {
			http.Error(w, "Username and/or password do not match", http.StatusForbidden)
			return
		}
		// create session
		sID := uuid.NewV4()
		c := &http.Cookie{
			Name:  "session",
			Value: sID.String(),
		}
		c.MaxAge = session.Length
		http.SetCookie(w, c)
		session.Sessions[c.Value] = models.Session{un, time.Now()}
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}
	session.Show() // for demonstration purposes
	c.tpl.ExecuteTemplate(w, "login.gohtml", u)
}

// Logout exported
func (c Controller) Logout(w http.ResponseWriter, req *http.Request) {
	if !session.AlreadyLoggedIn(w, req) {
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}
	ck, _ := req.Cookie("session")
	// delete the session
	delete(session.Sessions, ck.Value)
	// remove the cookie
	ck = &http.Cookie{
		Name:   "session",
		Value:  "",
		MaxAge: -1,
	}
	http.SetCookie(w, ck)

	// clean up dbSessions
	if time.Now().Sub(session.LastCleaned) > (time.Second * 30) {
		go session.Clean()
	}

	http.Redirect(w, req, "/login", http.StatusSeeOther)
}
