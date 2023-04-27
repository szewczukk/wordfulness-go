package main

import (
	"database/sql"
	"html/template"
	"log"
	"net/http"
	"wordfulness/core"
	"wordfulness/services"
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

	templates := map[string]*template.Template{
		"HomePage": template.Must(template.ParseFiles(
			"templates/layout.html",
			"templates/homepage.html",
		)),
		"DetailedCourse": template.Must(template.ParseFiles(
			"templates/layout.html",
			"templates/course.html",
		)),
	}

	router := core.NewRouter()
	coursesController := services.NewCoursesService(storage, templates)

	router.Get("/", coursesController.HomePage)
	router.Post("/", coursesController.CreateCourse)
	router.Get("/courses", coursesController.DetailedCourse)
	router.Get("/delete-course", coursesController.DeleteCourse)

	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	http.Handle("/", router)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
