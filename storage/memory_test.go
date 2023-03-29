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
