package storage

import (
	"database/sql"
	"log"
	"wordfulness/types"
)

type SequelStorage struct {
	Db *sql.DB
}

func (storage *SequelStorage) Initialize() {
	query := `
		CREATE TABLE IF NOT EXISTS courses (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			name VARCHAR(20) UNIQUE
		);
	`

	_, err := storage.Db.Exec(query)
	if err != nil {
		log.Fatal(err)
	}
}

func (storage *SequelStorage) GetAllCourses() ([]*types.Course, error) {
	var courses []*types.Course

	rows, err := storage.Db.Query("SELECT id, name FROM courses")

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		course := new(types.Course)

		err = rows.Scan(&course.Id, &course.Name)
		if err != nil {
			return nil, err
		}

		courses = append(courses, course)
	}

	return courses, nil
}

func (storage *SequelStorage) GetCourse(id int64) (*types.Course, error) {
	row := storage.Db.QueryRow("SELECT id, name FROM courses WHERE id = ?", id)

	course := new(types.Course)
	err := row.Scan(&course.Id, &course.Name)
	if err != nil {
		return nil, err
	}

	return course, nil
}

func (storage *SequelStorage) CreateCourse(name string) error {
	_, err := storage.Db.Exec("INSERT INTO courses (name) VALUES (?)", name)

	if err != nil {
		return err
	}

	return nil
}

func (storage *SequelStorage) DeleteCourse(id int64) error {
	_, err := storage.Db.Exec("DELETE FROM courses WHERE id = ?", id)

	if err != nil {
		return err
	}

	return nil
}
