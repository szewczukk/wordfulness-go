package main

import (
	"database/sql"
	"html/template"
	"log"
	"net/http"
	"wordfulness/core"
	"wordfulness/routes"
	"wordfulness/storage"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "database.sqlite")
	if err != nil {
		log.Fatal(err)
	}

	storage := &storage.SequelStorage{Db: db}
	storage.Initialize()

	homePageTemplate := template.Must(template.ParseFiles(
		"templates/layout.html",
		"templates/homepage.html",
	))
	courseTemplate := template.Must(template.ParseFiles(
		"templates/layout.html",
		"templates/course.html",
	))

	router := core.NewRouter()

	router.Get("/", routes.HomePage(storage, homePageTemplate))
	router.Post("/", routes.CreateCourse(storage, homePageTemplate))
	router.Get("/courses", routes.DetailedCourse(storage, courseTemplate))
	router.Get("/delete-course", routes.DeleteCourse(storage))

	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	http.Handle("/", router)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
