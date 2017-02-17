package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"net/url"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB
var err error
var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))

}

func main() {
	db, err = sql.Open("mysql", "root:Password@tcp(localhost:3306)/amigos?charset=utf8")
	check(err)
	defer db.Close()

	err = db.Ping()
	check(err)

	http.HandleFunc("/", index)
	http.HandleFunc("/ping", ping)
	http.HandleFunc("/instance", instance)
	http.HandleFunc("/amigos", amigos)
	http.HandleFunc("/create", create)
	http.HandleFunc("/insert", insert)
	http.HandleFunc("/read", read)
	http.HandleFunc("/update", update)
	http.HandleFunc("/delete", delete)
	http.HandleFunc("/drop", drop)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	err := http.ListenAndServe(":80", nil)
	check(err)
}

func index(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello from CAM Server 1")
}

func ping(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "OK")
}

func instance(w http.ResponseWriter, r *http.Request) {
	data := struct {
		Method        string
		URL           *url.URL
		Submissions   url.Values
		Header        http.Header
		Host          string
		ContentLength int64
	}{
		r.Method,
		r.URL,
		r.Form,
		r.Header,
		r.Host,
		r.ContentLength,
	}
	tpl.ExecuteTemplate(w, "index.gohtml", data)
}

func amigos(w http.ResponseWriter, r *http.Request) {
	rows, err := db.Query(`SELECT name FROM amigos;`)
	check(err)

	var s, name string
	s = "RETRIEVED RECORDS:\n"

	for rows.Next() {
		err = rows.Scan(&name)
		check(err)
		s += name + "\n"
	}
	fmt.Fprintln(w, s)
}

func read(w http.ResponseWriter, r *http.Request) {
	rows, err := db.Query(`SELECT * FROM customer;`)
	check(err)

	var name string
	for rows.Next() {
		err = rows.Scan(&name)
		check(err)
		fmt.Fprintln(w, "RETRIEVED RECORD:", name)
	}
}

func create(w http.ResponseWriter, r *http.Request) {
	stmt, err := db.Prepare(`CREATE TABLE customer(name VARCHAR(20));`)
	check(err)

	rows, err := stmt.Exec()
	check(err)

	n, err := rows.RowsAffected()
	check(err)

	fmt.Fprintln(w, "CREATED TABLE CUSTOMER", n)
}

func insert(w http.ResponseWriter, r *http.Request) {
	stmt, err := db.Prepare(`INSERT INTO customer VALUES("James");`)
	check(err)

	rows, err := stmt.Exec()
	check(err)

	n, err := rows.RowsAffected()
	check(err)

	fmt.Fprintln(w, "INSERTED RECORD", n)
}

func update(w http.ResponseWriter, r *http.Request) {
	stmt, err := db.Prepare(`UPDATE customer SET name="Jimmy" WHERE name="James";`)
	check(err)

	rows, err := stmt.Exec()
	check(err)

	n, err := rows.RowsAffected()
	check(err)

	fmt.Fprintln(w, "UPDATED RECORD", n)
}
func delete(w http.ResponseWriter, r *http.Request) {
	stmt, err := db.Prepare(`DELETE FROM customer WHERE name="Jimmy";`)
	check(err)

	rows, err := stmt.Exec()
	check(err)

	n, err := rows.RowsAffected()
	check(err)

	fmt.Fprintln(w, "DELETED RECORD", n)
}
func drop(w http.ResponseWriter, r *http.Request) {
	stmt, err := db.Prepare(`DROP TABLE customer;`)
	check(err)

	_, err = stmt.Exec()
	check(err)
	fmt.Fprintln(w, "DROPPED TABLE customer")
}

func check(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
