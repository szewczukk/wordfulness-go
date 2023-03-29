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

func (s *MemoryStorage) GetCourse(id int64) (*types.Course, error) {
	for _, course := range s.courses {
		if course.Id == int(id) {
			return course, nil
		}
	}

	return nil, errors.New("not found")
}

func (s *MemoryStorage) CreateCourse(name string) error {
	s.courses = append(s.courses, &types.Course{Id: len(s.courses), Name: name})
	return nil
}
