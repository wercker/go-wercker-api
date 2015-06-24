package wercker

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

// APIError represents a wercker error.
type APIError struct {
	Message    string `json:"message"`
	StatusCode int    `json:"statusCode"`
}

// Error returns the message and status code.
func (e *APIError) Error() string {
	return fmt.Sprintf("wercker-api: %s (status code: %d)", e.Message, e.StatusCode)
}

// parseError will check if res.Body contains a wercker generated error and
// return that, otherwise it will return a generic message based on statuscode.
func parseError(res *http.Response) error {
	// Check if the Body contains a wercker JSON error.
	if res.ContentLength > 0 {
		contentType := strings.Trim(res.Header.Get("Content-Type"), " ")

		if strings.HasPrefix(contentType, "application/json") {
			buf, err := ioutil.ReadAll(res.Body)
			if err != nil {
				goto generic
			}
			defer res.Body.Close()

			var payload *APIError
			err = json.Unmarshal(buf, &payload)
			if err == nil && payload.Message != "" && payload.StatusCode != 0 {
				return payload
			}
		}
	}

generic:
	var message string
	switch res.StatusCode {
	case 401:
		message = "authentication required"
	case 403:
		message = "not authorized to access this resource"
	case 404:
		message = "resource not found"
	default:
		message = "unknown error"
	}

	return &APIError{
		Message:    message,
		StatusCode: res.StatusCode,
	}
}
