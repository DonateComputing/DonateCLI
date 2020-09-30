package api

// ResponseStruct is generic struct that api server responds with on resource edit, create
type ResponseStruct struct {
	Success   bool   `json:"success"`
	Message   string `json:"message"`
	AlteredID string `json:"alteredId"`
}
