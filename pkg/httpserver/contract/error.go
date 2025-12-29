package httpservercontract

// HTTPError represents a standardized HTTP error that can be serialized as JSON.
// Follows the specified format with error code, message, and causes.
//
// JSON format:
//
//	{
//	  "code": "ANY0001",
//	  "message": "Error message",
//	  "cause": ["id invalid", "name is required"]
//	}

// TODO: maybe change this interface to be a struct
type HTTPError interface {
	error
	// Status returns the HTTP status code associated with the error
	Status() int

	// Code returns the error code (e.g., "ANY0001")
	Code() string

	// Message returns the descriptive error message
	Message() string

	// Causes returns the list of specific error causes/messages
	// Can be empty if there are no detailed causes
	Causes() []string
}

const (
	HttpServerUnexpectedError = "HTTPSERVER0001"
)
