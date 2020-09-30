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
		t.Fatalf("PostJob('%v') error '%v'", job, err)
	}
}

func TestGetJob(t *testing.T) {
	ref := JobRefStruct{auth.Username, "TESTJOB"}
	job, err := GetJob(ref)
	if err != nil {
		t.Fatalf("GetJob('%v') error '%v'", ref, err)
	}
	if job.Title != ref.Title || job.Author != ref.User {
		t.Fatalf("GetJob('%v') returned bad job '%v'", ref, job)
	}
}
