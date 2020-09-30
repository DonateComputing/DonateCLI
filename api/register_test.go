package api

import (
	"testing"
)

var auth = AuthStruct{
	Username: "miko",
	Password: "pass",
}

func TestPostRegister(t *testing.T) {
	err := PostRegister(auth)
	if err != nil {
		if err.Error() == "Error: user miko already exists" {
			return
		}
		t.Fatalf("PostRegister() error '%s'", err)
	}
}
