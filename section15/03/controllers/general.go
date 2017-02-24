package controllers

import (
	"html/template"
	"net/http"

	session "github.com/ramirezra/golang-web/section15/03/sessions"
)

// Controller exported
type Controller struct {
	tpl *template.Template
}

// NewController exported
func NewController(t *template.Template) *Controller {
	return &Controller{t}
}

// Index exported
func (c Controller) Index(w http.ResponseWriter, req *http.Request) {
	u := session.GetUser(w, req)
	session.Show() // for demonstration purposes
	c.tpl.ExecuteTemplate(w, "index.gohtml", u)
}

// Bar exported
func (c Controller) Bar(w http.ResponseWriter, req *http.Request) {
	u := session.GetUser(w, req)
	if !session.AlreadyLoggedIn(w, req) {
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}
	if u.Role != "007" {
		http.Error(w, "You must be 007 to enter the bar", http.StatusForbidden)
		return
	}
	session.Show() // for demonstration purposes
	c.tpl.ExecuteTemplate(w, "bar.gohtml", u)
}
