package storage

import (
	"database/sql"
	"log"
	"wordfulness/types"
)

type SequelStorage struct {
	db *sql.DB
}

func NewSequelStorage(db *sql.DB) *SequelStorage {
	return &SequelStorage{
		db: db,
	}
}

func (s *SequelStorage) Initialize() {
	query := `
		CREATE TABLE IF NOT EXISTS courses (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			name VARCHAR(20) UNIQUE
		);

		CREATE TABLE IF NOT EXISTS users (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			username VARCHAR(20) UNIQUE,
			password CHAR(60)
		);
	`

	_, err := s.db.Exec(query)
	if err != nil {
		log.Fatal(err)
	}
}

func (s *SequelStorage) GetAllCourses() ([]*types.Course, error) {
	var courses []*types.Course

	rows, err := s.db.Query("SELECT id, name FROM courses")
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

func (s *SequelStorage) GetCourse(id int) (*types.Course, error) {
	row := s.db.QueryRow("SELECT id, name FROM courses WHERE id = ?", id)

	course := new(types.Course)
	err := row.Scan(&course.Id, &course.Name)
	if err != nil {
		return nil, err
	}

	return course, nil
}

func (s *SequelStorage) CreateCourse(name string) error {
	_, err := s.db.Exec("INSERT INTO courses (name) VALUES (?)", name)
	if err != nil {
		return err
	}

	return nil
}

func (s *SequelStorage) DeleteCourse(id int) error {
	_, err := s.db.Exec("DELETE FROM courses WHERE id = ?", id)
	if err != nil {
		return err
	}

	return nil
}

func (s *SequelStorage) UpdateCourse(id int, name string) error {
	_, err := s.db.Exec("UPDATE courses SET name = ? WHERE id = ?", name, id)
	if err != nil {
		return err
	}

	return nil
}

func (s *SequelStorage) CreateUser(name string, password string) error {
	_, err := s.db.Exec("INSERT INTO users (username, password) VALUES (?, ?)", name, password)
	if err != nil {
		return err
	}

	return nil
}

func (s *SequelStorage) GetUserByUserName(username string) (*types.User, error) {
	row := s.db.QueryRow("SELECT id, username, password FROM users WHERE username = ?", username)

	user := new(types.User)
	err := row.Scan(&user.Id, &user.Username, &user.Password)
	if err != nil {
		return nil, err
	}

	return user, nil
}
