package errors

import (
	"fmt"
	"net/http"
)

// HTTPError is a custom error for http requests
type HTTPError struct {
	StatusCode int
	Message    string
}

func (err *HTTPError) Error() string {
	return fmt.Sprintf("[%d] => %s", err.StatusCode, err.Message)
}

// NewHTTPError is a function for generating a NewHTTPError
func NewHTTPError(statusCode int, messagesOverride map[string]interface{}) (err *HTTPError) {
	err = &HTTPError{StatusCode: statusCode}

	switch statusCode {
	case http.StatusNotFound:
		err.Message = "Record Not Found"
	case http.StatusBadRequest:
		err.Message = "Bad Request"
	case http.StatusInternalServerError:
		err.Message = "Service Unavailable"
	default:
		fmt.Println(statusCode)
	}
	
	return
}