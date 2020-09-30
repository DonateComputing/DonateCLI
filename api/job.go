package api

// JobStruct is struct returned by api when querying jobs
type JobStruct struct {
	Title          string `json:"title"`
	Description    string `json:"description"`
	OriginalImage  string `json:"originalImage"`
	CompletedImage string `json:"completedImage"`
	Timestamp      int64  `json:"timestamp"`
	Author         string `json:"author"`
	Runner         string `json:"runner"`
}

// JobMapStruct is struct returned by api when querying a user's jobs
type JobMapStruct map[string]JobStruct
