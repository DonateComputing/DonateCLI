package api

import (
	"fmt"
	"net/http"
)

// ResponseStruct is generic struct that api server responds with on resource edit, create
type ResponseStruct struct {
	Success   bool   `json:"success"`
	Message   string `json:"message"`
	AlteredID string `json:"alteredId"`
}

func makeErrorFromResponse(res *http.Response, r ResponseStruct) error {
	return fmt.Errorf("`%s` rejected (%d), '%s'", res.Request.URL, res.StatusCode, r.Message)
}
