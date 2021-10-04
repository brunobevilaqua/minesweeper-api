package errors

import "net/http"

type ApiError struct {
	Message    string `json:"-"`
	StatusCode int    `json:"type"`
	Type       string `json:"message,omitempty"`
}

type ErrorType int

const (
	InvalidUserName ErrorType = iota
	InvalidParameter
	NoRecordsFound
	ErrorSavingEntity
)

func NewApiError(e ErrorType) *ApiError {
	return e.getErrorByType()
}

func (et ErrorType) getErrorByType() *ApiError {
	switch et {
	case InvalidUserName:
		message := "User name cannot be blank."
		statusCode := http.StatusBadRequest
		typee := "invalid_json"
		return &ApiError{Message: message, Type: typee, StatusCode: statusCode}
	case InvalidParameter:
		message := "Invalid Parameter."
		statusCode := http.StatusBadRequest
		typee := "invalid_json"
		return &ApiError{Message: message, Type: typee, StatusCode: statusCode}
	case NoRecordsFound:
		message := "No Records Found."
		statusCode := http.StatusNoContent
		typee := "database"
		return &ApiError{Message: message, Type: typee, StatusCode: statusCode}
	case ErrorSavingEntity:
		message := "Error when saving Entity in database."
		statusCode := http.StatusInternalServerError
		typee := "database"
		return &ApiError{Message: message, Type: typee, StatusCode: statusCode}
	}

	return &ApiError{Message: "An unknown error occured.", StatusCode: http.StatusInternalServerError, Type: "server_error"}
}
