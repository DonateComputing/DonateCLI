package api

import (
	"testing"
)

func TestGetUser(t *testing.T) {
	ensureUserExists(auth)
	u, err := GetUser(auth)
	if err != nil {
		t.Fatalf("GetUser() error '%v'", err)
	}
	if u.Username != auth.Username {
		t.Fatalf("GetUser() returned bad id '%s' instead of '%s'", u.Username, auth.Username)
	}
}

func TestUpdatePassword(t *testing.T) {
	ensureUserExists(auth)
	err := UpdatePassword("TEMP", auth)
	if err != nil {
		t.Fatalf("UpdatePassword('TEMP') error '%v'", err)
	}

	newAuth := AuthStruct{
		Username: auth.Username,
		Password: "TEMP",
	}
	err = UpdatePassword(auth.Password, newAuth)
	if err != nil {
		t.Fatalf("UpdatePassword('TEMP') secondary error '%v'", err)
	}
}

func TestDeleteUser(t *testing.T) {
	ensureUserExists(auth)
	err := DeleteUser(auth)
	if err != nil {
		t.Fatalf("DeleteUser('%v') error '%v'", auth, err)
	}

	err = RegisterUser(auth)
	if err != nil {
		t.Fatalf("DeleteUser('%v') cleanup failed: '%v'", auth, err)
	}
}
