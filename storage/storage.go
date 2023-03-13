package storage

type IStorage interface {
	GetAllCourses() []Course
	CreateCourse(name string)
}
