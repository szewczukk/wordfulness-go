package storage

type Course struct {
	Id   int
	Name string
}

type MemoryStorage struct {
	Courses []Course
}

func (storage *MemoryStorage) GetAllCourses() []Course {
	return storage.Courses
}

func (storage *MemoryStorage) CreateCourse(name string) {
	course := &Course{Name: name}

	storage.Courses = append(storage.Courses, *course)
}
