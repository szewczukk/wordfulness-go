package storage_test

import (
	"testing"
	"wordfulness/storage"
	"wordfulness/types"
)

func TestCreateUser(t *testing.T) {
	storage := storage.NewUserMemoryStorage([]*types.User{})

	err := storage.CreateUser("jbytnar", "zaq1@WSX")

	if err != nil {
		t.Errorf("Error occurred, got: %v", err.Error())
	}
}

func TestCreateUserDuplicate(t *testing.T) {
	storage := storage.NewUserMemoryStorage([]*types.User{{Id: 0, Username: "jbytnar", Password: ""}})

	err := storage.CreateUser("jbytnar", "zaq1@WSX")

	if err.Error() != "duplicate" {
		t.Errorf("Incorrect error occurred, got: %v", err.Error())
	}
}

func TestGetUserByUserName(t *testing.T) {
	storage := storage.NewUserMemoryStorage([]*types.User{{Id: 0, Username: "jbytnar", Password: ""}})

	user, err := storage.GetUserByUserName("jbytnar")

	if err != nil {
		t.Errorf("Error occurred, got: %v", err.Error())
	}

	if user.Username != "jbytnar" {
		t.Errorf("Error username, got: %v", user.Username)
	}
}

func TestGetUserByUserNameNotFound(t *testing.T) {
	storage := storage.NewUserMemoryStorage([]*types.User{})

	_, err := storage.GetUserByUserName("jbytnar")

	if err.Error() != "not found" {
		t.Errorf("Incorrect error occurred, got: %v", err.Error())
	}
}

func TestGetUserById(t *testing.T) {
	storage := storage.NewUserMemoryStorage([]*types.User{{Id: 0, Username: "jbytnar", Password: ""}})

	user, err := storage.GetUserById(0)

	if err != nil {
		t.Errorf("Error occurred, got: %v", err.Error())
	}

	if user.Username != "jbytnar" {
		t.Errorf("Error username, got: %v", user.Username)
	}
}

func TestGetUserByIdNotFound(t *testing.T) {
	storage := storage.NewUserMemoryStorage([]*types.User{})

	_, err := storage.GetUserById(0)

	if err.Error() != "not found" {
		t.Errorf("Incorrect error occurred, got: %v", err.Error())
	}
}
