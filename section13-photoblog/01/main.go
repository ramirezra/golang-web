package main

import (
	"html/template"
	"net/http"

	"github.com/satori/go.uuid"
)

var tpl *template.Template

// var dbUsers = map[string]user{}      // userID, user
// var dbSessions = map[string]string{} // session ID, user ID

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {
	http.HandleFunc("/", index)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	// u := getUser(w, r)
	c := getCookie(w, r)
	tpl.ExecuteTemplate(w, "index.gohtml", c)
}

// Session and Cookie Setting
// type user struct {
// 	UserName string
// 	Password []byte
// 	First    string
// 	Last     string
// }

func getCookie(w http.ResponseWriter, r *http.Request) *http.Cookie {
	c, err := r.Cookie("session")
	if err != nil {
		sID := uuid.NewV4()
		c = &http.Cookie{
			Name:  "session",
			Value: sID.String(),
		}
	}
	http.SetCookie(w, c)

	// Get user, if already exists.
	// var u user
	// if un, ok := dbSessions[c.Value]; ok {
	// 	u = dbUsers[un]
	// }
	return c
}
