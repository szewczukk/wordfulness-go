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
		t.Errorf("Incorrect occurred, got: %v", err.Error())
	}
}
