package api

import "testing"

func TestGetJobs(t *testing.T) {
	_, err := GetJobs()
	if err != nil {
		t.Fatalf("GetJobs() error '%v'", err)
	}
}

func TestPostJob(t *testing.T) {
	job := PostJobStruct{
		Title:       "TESTJOB",
		Description: "TESTDESCRIPTION",
		Image:       "TESTIMAGE",
	}
	err := PostJob(job, auth)
	if err != nil {
		t.Fatalf("PostJob() error '%v'", err)
	}
}
