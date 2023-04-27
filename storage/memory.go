package storage

import (
	"errors"
	"wordfulness/types"
)

type MemoryStorage struct {
	courses []*types.Course
}

func NewMemoryStorage(courses []*types.Course) *MemoryStorage {
	return &MemoryStorage{
		courses: courses,
	}
}

func (s *MemoryStorage) GetAllCourses() ([]*types.Course, error) {
	return s.courses, nil
}

func (s *MemoryStorage) GetCourse(id int) (*types.Course, error) {
	for _, course := range s.courses {
		if course.Id == id {
			return course, nil
		}
	}

	return nil, errors.New("not found")
}

func (s *MemoryStorage) CreateCourse(name string) error {
	for _, course := range s.courses {
		if course.Name == name {
			return errors.New("duplicate")
		}
	}
	s.courses = append(s.courses, &types.Course{Id: len(s.courses), Name: name})
	return nil
}

func (s *MemoryStorage) DeleteCourse(id int) error {
	filteredCourses := []*types.Course{}
	courseExists := false

	for _, course := range s.courses {
		if course.Id == id {
			filteredCourses = append(filteredCourses, course)
			courseExists = true
		}
	}

	if !courseExists {
		return errors.New("not found")
	}

	s.courses = filteredCourses

	return nil
}
