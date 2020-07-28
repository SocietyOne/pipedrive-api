package pipedrive

import (
	"fmt"
	"net/http"
)

// RateLimitError occurs when Pipedrive returns 403 Forbidden response with a rate limit
// remaining value of 0.
type RateLimitError struct {
	Rate     Rate
	Response *http.Response
	Message  string
}

func (e *RateLimitError) Error() string {
	return fmt.Sprintf("%v %v: %d %v",
		e.Response.Request.Method, e.Response.Request.URL,
		e.Response.StatusCode, e.Message)
}

// ResponseError sometimes error is a custom struct instead of string
// TODO: expand on different error handling
type ResponseError string

func (e *ResponseError) UnmarshalJSON(b []byte) error {
	*e = ResponseError(string(b))
	return nil
}

// ErrorResponse reports one or more errors caused by an API request.
type ErrorResponse struct {
	Response *http.Response
	Message  ResponseError `json:"error"`
}

func (e *ErrorResponse) Error() string {
	return fmt.Sprintf("%v: %d %v",
		e.Response.Request.Method, e.Response.StatusCode, string(e.Message))
}
