package main

import (
	"database/sql"
	"html/template"
	"log"
	"net/http"
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

	router := routes.NewRouter()

	router.GET("/", routes.HomePage(storage, homePageTemplate))
	router.POST("/", routes.CreateCourse(storage, homePageTemplate))
	router.GET("/courses", routes.DetailedCourse(storage, courseTemplate))
	router.GET("/delete-course", routes.DeleteCourse(storage))

	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	http.Handle("/", router)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
