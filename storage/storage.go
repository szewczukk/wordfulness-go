package storage

import "wordfulness/types"

type IStorage interface {
	GetAllCourses() []types.Course
	CreateCourse(name string) error
}
