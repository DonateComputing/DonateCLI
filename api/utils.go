package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func doRequest(method string, url string, data interface{}) (*http.Response, error) {
	// marshal data (if exists)
	var bytesBuffer *bytes.Buffer
	if data != nil {
		reqBytes, err := json.Marshal(data)
		if err != nil {
			return &http.Response{}, err
		}
		bytesBuffer = bytes.NewBuffer(reqBytes)
	}
	// set up request
	req, err := http.NewRequest(method, url, bytesBuffer)
	if err != nil {
		return &http.Response{}, err
	}
	if data != nil {
		req.Header.Add("Content-Type", "application/json")
	}
	// do request
	return client.Do(req)
}

func parseResponseBody(response *http.Response, data interface{}) error {
	// check headers
	if response.Header.Get("Content-Type") != "application/json" {
		return fmt.Errorf("content type indicates not json: %s", response.Header.Get("Content-Type"))
	}
	// read bytes
	defer response.Body.Close()
	bodyBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return err
	}
	// unmarshal bytes
	err = json.Unmarshal(bodyBytes, data)
	if err != nil {
		return fmt.Errorf("failed to unmarshal response into json: %v\n\t%s", err, bodyBytes)
	}

	return nil
}
