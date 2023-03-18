package storage

import "wordfulness/types"

type IStorage interface {
	GetAllCourses() ([]*types.Course, error)
	CreateCourse(name string) error
}
