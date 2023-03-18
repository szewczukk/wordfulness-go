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
	multipleCoursesTemplate := template.Must(template.ParseFiles("templates/homepage.html"))
	singleCourseTemplate := template.Must(template.ParseFiles("templates/course.html"))

	http.HandleFunc("/", api.GetMultipleCourses(storage, multipleCoursesTemplate))
	http.HandleFunc("/courses", api.GetSingleCourse(storage, singleCourseTemplate))

	log.Fatal(http.ListenAndServe(":8080", nil))
}
