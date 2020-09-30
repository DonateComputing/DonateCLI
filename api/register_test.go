package api

import (
	"testing"
)

var auth = AuthStruct{
	Username: "miko",
	Password: "pass",
}

func TestRegisterUser(t *testing.T) {
	err := RegisterUser(auth)
	if err != nil {
		if err.Error() == "Error: user miko already exists" {
			return
		}
		t.Fatalf("PostRegister('%v') error '%s'", auth, err)
	}
}
