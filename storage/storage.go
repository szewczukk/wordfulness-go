package storage

import "wordfulness/types"

type IStorage interface {
	GetAllCourses() ([]*types.Course, error)
	GetCourse(id int) (*types.Course, error)
	CreateCourse(name string) error
	DeleteCourse(id int) error
}
