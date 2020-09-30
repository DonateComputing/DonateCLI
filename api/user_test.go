package api

import (
	"testing"
)

func TestGetUser(t *testing.T) {
	u, err := GetUser(auth)
	if err != nil {
		t.Fatalf("GetUser() error '%v'", err)
	}
	if u.Username != auth.Username {
		t.Fatalf("GetUser() returned bad id '%s' instead of '%s'", u.Username, auth.Username)
	}
}

func TestUpdatePassword(t *testing.T) {
	err := UpdatePassword("TEMP", auth)
	if err != nil {
		t.Fatalf("UpdatePassword() error '%v'", err)
	}

	newAuth := AuthStruct{
		Username: auth.Username,
		Password: "TEMP",
	}
	err = UpdatePassword(auth.Password, newAuth)
	if err != nil {
		t.Fatalf("UpdatePassword() secondary error '%v'", err)
	}
}

func TestDeleteUser(t *testing.T) {
	err := DeleteUser(auth)
	if err != nil {
		t.Fatalf("DeleteUser() error '%v'", err)
	}
}
