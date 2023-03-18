package main

import (
	"database/sql"
	"html/template"
	"log"
	"net/http"
	"wordfulness/api"
	"wordfulness/storage"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "database.sqlite")

	if err != nil {
		log.Fatal(err)
	}

	storage := &storage.SequelStorage{Db: db}
	tmpl := template.Must(template.ParseFiles("templates/homepage.html"))

	http.HandleFunc("/", api.GetCourses(storage, tmpl))

	log.Fatal(http.ListenAndServe(":8080", nil))
}
