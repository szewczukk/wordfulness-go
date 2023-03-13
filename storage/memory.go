package storage

import "errors"

type Course struct {
	Id   int
	Name string
}

type MemoryStorage struct {
	Courses []Course
}

func (storage *MemoryStorage) GetAllCourses() []Course {
	return storage.Courses
}

func (storage *MemoryStorage) CreateCourse(name string) error {
	course := &Course{Name: name}

	for _, course := range storage.Courses {
		if course.Name == name {
			return errors.New("Duplicate")
		}
	}

	storage.Courses = append(storage.Courses, *course)

	return nil
}
