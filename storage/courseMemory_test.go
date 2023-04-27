package storage_test

import (
	"reflect"
	"testing"
	"wordfulness/storage"
	"wordfulness/types"
)

func TestGetAllCoursesReturnsEmptySlice(t *testing.T) {
	storage := storage.NewCourseMemoryStorage([]*types.Course{})

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
	storage := storage.NewCourseMemoryStorage(
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
	storage := storage.NewCourseMemoryStorage(
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
	storage := storage.NewCourseMemoryStorage([]*types.Course{})

	course, err := storage.GetCourse(0)

	if course != nil {
		t.Errorf("Returned course: %v", course)
	}

	if err.Error() != "not found" {
		t.Errorf("Invalid error, got: %v", err)
	}
}

func TestCreateCourseReturnsNil(t *testing.T) {
	storage := storage.NewCourseMemoryStorage([]*types.Course{{}})

	err := storage.CreateCourse("Spanish")

	if err != nil {
		t.Errorf("Error occurred: %v", err)
	}
}

func TestCreateCourseReturnsError(t *testing.T) {
	storage := storage.NewCourseMemoryStorage([]*types.Course{{Id: 0, Name: "Spanish"}})

	err := storage.CreateCourse("Spanish")

	if err == nil {
		t.Error("Error didn't occurr")
	}

	if err.Error() != "duplicate" {
		t.Errorf("Wrong error returned, got: %v", err)
	}
}

func TestDeleteCourseReturnsNil(t *testing.T) {
	storage := storage.NewCourseMemoryStorage([]*types.Course{{Id: 0, Name: "Spanish"}})

	err := storage.DeleteCourse(0)

	if err != nil {
		t.Errorf("Error occurred: %v", err)
	}
}

func TestDeleteCourseReturnsNotFoundError(t *testing.T) {
	storage := storage.NewCourseMemoryStorage([]*types.Course{})

	err := storage.DeleteCourse(0)

	if err == nil {
		t.Error("Error didn't occurr")
	}

	if err.Error() != "not found" {
		t.Errorf("Wrong error returned, got: %v", err)
	}
}

func TestUpdateCourse(t *testing.T) {
	storage := storage.NewCourseMemoryStorage([]*types.Course{{Id: 0, Name: "Spanish"}})

	err := storage.UpdateCourse(0, "German")

	if err != nil {
		t.Errorf("Error: %v", err)
	}
}

func TestUpdateCourseReturnsDuplicateError(t *testing.T) {
	storage := storage.NewCourseMemoryStorage([]*types.Course{{Id: 0, Name: "German"}, {Id: 1, Name: "Spanish"}})

	err := storage.UpdateCourse(1, "German")

	if err == nil {
		t.Error("Error didn't occurr")
	}

	if err.Error() != "duplicate" {
		t.Errorf("Wrong error returned, got: %v", err)
	}
}

func TestUpdateCourseReturnsNotFoundError(t *testing.T) {
	storage := storage.NewCourseMemoryStorage([]*types.Course{})

	err := storage.UpdateCourse(0, "German")

	if err == nil {
		t.Error("Error didn't occurr")
	}

	if err.Error() != "not found" {
		t.Errorf("Wrong error returned, got: %v", err)
	}
}
