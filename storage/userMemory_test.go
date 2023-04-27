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

func TestGetUser(t *testing.T) {
	storage := storage.NewUserMemoryStorage([]*types.User{{Id: 0, Username: "jbytnar", Password: ""}})

	user, err := storage.GetUser(0)

	if err != nil {
		t.Errorf("Error occurred, got: %v", err.Error())
	}

	if user.Username != "jbytnar" {
		t.Errorf("Error username, got: %v", user.Username)
	}
}

func TestGetUserNotFound(t *testing.T) {
	storage := storage.NewUserMemoryStorage([]*types.User{})

	_, err := storage.GetUser(0)

	if err.Error() != "not found" {
		t.Errorf("Incorrect error occurred, got: %v", err.Error())
	}
}
