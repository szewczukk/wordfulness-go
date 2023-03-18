package storage

import (
	"errors"
	"wordfulness/types"
)

type MemoryStorage struct {
	Courses []*types.Course
}

func (storage *MemoryStorage) GetAllCourses() ([]*types.Course, error) {
	return storage.Courses, nil
}

func (storage *MemoryStorage) CreateCourse(name string) error {
	course := &types.Course{Name: name}

	for _, course := range storage.Courses {
		if course.Name == name {
			return errors.New("Duplicate")
		}
	}

	storage.Courses = append(storage.Courses, course)

	return nil
}
