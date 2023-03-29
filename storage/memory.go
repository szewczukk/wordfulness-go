package storage

import (
	"errors"
	"wordfulness/types"
)

type MemoryStorage struct {
	courses []*types.Course
	nextId  int64
}

func NewMemoryStorage(courses []*types.Course, nextId int64) *MemoryStorage {
	return &MemoryStorage{
		courses: courses,
		nextId:  nextId,
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
	s.courses = append(s.courses, &types.Course{Id: int(s.nextId), Name: name})
	s.nextId++
	return nil
}
