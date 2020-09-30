package api

import "testing"

func TestGetJobs(t *testing.T) {
	_, err := GetJobs()
	if err != nil {
		t.Fatalf("GetJobs() error '%v'", err)
	}
}
