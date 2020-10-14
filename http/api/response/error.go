package response

// Error to represent error
type Error struct {
	StatusCode int    `json:"-"`
	Code       string `json:"code"`
	Message    string `json:"message"`
}

// ServerError representing server side error
var ServerError Error = Error{
	StatusCode: 500,
	Code:       "server_error",
	Message:    "Server Error",
}

// NotFoundError representing server side error
var NotFoundError Error = Error{
	StatusCode: 404,
	Code:       "not_found",
	Message:    "Not Found",
}
