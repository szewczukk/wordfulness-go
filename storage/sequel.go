package storage

import (
	"database/sql"
	"wordfulness/types"
)

type SequelStorage struct {
	Db *sql.DB
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

func (storage *SequelStorage) CreateCourse(name string) error {
	_, err := storage.Db.Exec("INSERT INTO courses (name) VALUES (?)", name)

	if err != nil {
		return err
	}

	return nil
}
