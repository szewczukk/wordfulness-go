package storage

import "wordfulness/types"

type IStorage interface {
	GetAllCourses() ([]*types.Course, error)
	GetCourse(id int64) (*types.Course, error)
	CreateCourse(name string) error
	DeleteCourse(id int64) error
}
