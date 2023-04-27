package storage

import (
	"errors"
	"wordfulness/types"
)

type CourseMemoryStorage struct {
	courses []*types.Course
}

func NewCourseMemoryStorage(courses []*types.Course) *CourseMemoryStorage {
	return &CourseMemoryStorage{
		courses: courses,
	}
}

func (s *CourseMemoryStorage) GetAllCourses() ([]*types.Course, error) {
	return s.courses, nil
}

func (s *CourseMemoryStorage) GetCourse(id int) (*types.Course, error) {
	for _, course := range s.courses {
		if course.Id == id {
			return course, nil
		}
	}

	return nil, errors.New("not found")
}

func (s *CourseMemoryStorage) CreateCourse(name string) error {
	for _, course := range s.courses {
		if course.Name == name {
			return errors.New("duplicate")
		}
	}
	s.courses = append(s.courses, &types.Course{Id: len(s.courses), Name: name})
	return nil
}

func (s *CourseMemoryStorage) DeleteCourse(id int) error {
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

func (s *CourseMemoryStorage) UpdateCourse(id int, name string) error {
	isChanged := false
	for _, course := range s.courses {
		if course.Name == name {
			return errors.New("duplicate")
		}

		if course.Id == id {
			course.Name = name
			isChanged = true
		}
	}

	if !isChanged {
		return errors.New("not found")
	}

	return nil
}
