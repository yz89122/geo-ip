package response

import (
	"github.com/gin-gonic/gin"
)

// Response object
type Response struct {
	result     interface{}
	success    bool
	errors     []*Error
	statusCode int
}

// New create a instance of Response
func New() *Response {
	response := Response{
		result:     nil,
		success:    true,
		errors:     []*Error{},
		statusCode: 0,
	}

	return &response
}

// SetResult set the result to the response object
func (r *Response) SetResult(result interface{}) {
	r.result = result
}

// Respond to client
func (r *Response) Respond(c *gin.Context) {
	response := struct {
		Result  interface{} `json:"result"`
		Success bool        `json:"success"`
		Errors  *[]*Error   `json:"errors"`
	}{
		Result:  r.result,
		Success: r.success,
		Errors:  &r.errors,
	}

	c.JSON(r.GetStatusCode(), &response)
}

// AppendError append an error to the response object
// If Error.StatusCode has been set, Response.StatusCode
// will be updated.
func (r *Response) AppendError(e Error) {
	r.errors = append(r.errors, &e)
	r.success = false
	if e.StatusCode != 0 {
		r.SetStatusCode(e.StatusCode)
	}
}

// SetStatusCode set the status code to the response object
func (r *Response) SetStatusCode(statusCode int) {
	r.statusCode = statusCode
}

// GetStatusCode return the status code of the response object
func (r *Response) GetStatusCode() int {
	if r.statusCode != 0 {
		return r.statusCode
	}

	if !r.success || len(r.errors) != 0 {
		return 500
	}

	return 200
}
