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

	storage := storage.NewSequelStorage(db)
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
		"UpdateCourse": template.Must(template.ParseFiles(
			"templates/layout.html",
			"templates/updateCourse.html",
		)),
		"CreateUser": template.Must(template.ParseFiles(
			"templates/layout.html",
			"templates/register.html",
		)),
	}

	router := core.NewRouter()
	coursesService := services.NewCoursesService(storage, templates)
	userService := services.NewUserService(storage, templates)

	router.Get("/", coursesService.HomePage)
	router.Post("/create-course", coursesService.CreateCourse)
	router.Post("/update-course", coursesService.UpdateCoursePOST)
	router.Get("/register", userService.CreateUserGet)
	router.Post("/register", userService.CreateUserPost)
	router.Get("/update-course", coursesService.UpdateCourseGET)
	router.Get("/courses", coursesService.DetailedCourse)
	router.Get("/delete-course", coursesService.DeleteCourse)

	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	http.Handle("/", router)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
