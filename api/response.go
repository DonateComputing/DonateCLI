package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// UpdateResponseStruct is generic struct that api server responds with on resource edit, create
type UpdateResponseStruct struct {
	Success   bool   `json:"success"`
	Message   string `json:"message"`
	AlteredID string `json:"alteredId"`
}

func parseUpdateResponseBody(res *http.Response) (*UpdateResponseStruct, error) {
	// read bytes
	defer res.Body.Close()
	bodyBytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return &UpdateResponseStruct{}, err
	}
	// unmarshal bytes
	var r UpdateResponseStruct
	err = json.Unmarshal(bodyBytes, &r)
	if err != nil {
		return &UpdateResponseStruct{}, fmt.Errorf("failed to unmarshal response into json: %v\n\t%s", err, bodyBytes)
	}

	return &r, err
}
