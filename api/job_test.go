package api

import "testing"

var job = PostJobStruct{
	Title:       "TESTJOB",
	Description: "TESTDESCRIPTION",
	Image:       "TESTIMAGE",
}

func TestGetJobs(t *testing.T) {
	_, err := GetJobs()
	if err != nil {
		t.Fatalf("GetJobs() error '%v'", err)
	}
}

func TestPostJob(t *testing.T) {
	err := PostJob(job, auth)
	if err != nil {
		t.Fatalf("PostJob('%v') error '%v'", job, err)
	}
}

func TestGetJob(t *testing.T) {
	ref := JobRefStruct{auth.Username, job.Title}
	job, err := GetJob(ref)
	if err != nil {
		t.Fatalf("GetJob('%v') error '%v'", ref, err)
	}
	if job.Title != ref.Title || job.Author != ref.User {
		t.Fatalf("GetJob('%v') returned bad job '%v'", ref, job)
	}
}

func TestDeleteJob(t *testing.T) {
	ref := JobRefStruct{auth.Username, job.Title}
	err := DeleteJob(ref, auth)
	if err != nil {
		t.Fatalf("DeleteJob('%v') error '%v'", ref, err)
	}

	err = PostJob(job, auth)
	if err != nil {
		t.Fatalf("DeleteJob('%v') cleanup failed '%v'", ref, err)
	}
}
