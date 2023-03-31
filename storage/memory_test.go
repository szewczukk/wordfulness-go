package storage_test

import (
	"reflect"
	"testing"
	"wordfulness/storage"
	"wordfulness/types"
)

func TestGetAllCoursesReturnsEmptySlice(t *testing.T) {
	storage := storage.NewMemoryStorage([]*types.Course{})

	courses, err := storage.GetAllCourses()

	if err != nil {
		t.Errorf("Error occurred: %v", err)
	}

	coursesCount := len(courses)
	if coursesCount != 0 {
		t.Errorf("Invalid coursesCount, expected 0 got %v", coursesCount)
	}
}

func TestGetAllCoursesReturnsTwoElements(t *testing.T) {
	storage := storage.NewMemoryStorage(
		[]*types.Course{
			{Id: 0, Name: "Spanish"}, {Id: 1, Name: "German"},
		},
	)

	courses, err := storage.GetAllCourses()

	if err != nil {
		t.Errorf("Error occured: %v", err)
	}

	coursesCount := len(courses)
	if coursesCount != 2 {
		t.Errorf("Invalid coursesCount, expected 2 got %v", coursesCount)
	}

	if !reflect.DeepEqual(courses, []*types.Course{{Id: 0, Name: "Spanish"}, {Id: 1, Name: "German"}}) {
		t.Errorf("Invalid courses, got: %v", courses)
	}
}

func TestGetCourseReturnsExistingCourse(t *testing.T) {
	storage := storage.NewMemoryStorage(
		[]*types.Course{
			{Id: 0, Name: "Spanish"}, {Id: 1, Name: "German"},
		},
	)

	course, err := storage.GetCourse(0)

	if err != nil {
		t.Errorf("Error occured: %v", err)
	}

	if !reflect.DeepEqual(course, &types.Course{Id: 0, Name: "Spanish"}) {
		t.Errorf("Invalid course, got: %v", course)
	}
}

func TestGetCourseReturnsError(t *testing.T) {
	storage := storage.NewMemoryStorage([]*types.Course{})

	course, err := storage.GetCourse(0)

	if course != nil {
		t.Errorf("Returned course: %v", course)
	}

	if err.Error() != "not found" {
		t.Errorf("Invalid error, got: %v", err)
	}
}
