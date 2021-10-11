package errors

import "net/http"

type ApiError struct {
	Message    string `json:"-"`
	StatusCode int    `json:"type"`
	Type       string `json:"message,omitempty"`
}

type ErrorType int

const (
	INVALID_USER_NAME_ERROR ErrorType = iota
	INVALID_PARAMETER_ERROR
	NO_RECORDS_FOUND_ERROR
	SAVING_GAME_ERROR
	SAVING_BOARD_ERROR
	INVALID_ACTION_ERROR
	CUSTOM_ERROR
	CELL_ALREADY_CLICKED_ERROR
	CELL_ALREADY_FLAGGED_ERROR
	LOST_GAME_ERROR
	UPDATE_CELL_STATUS_ERROR
	UPDATE_GAME_ERROR
	UPDATE_BOARD_ERROR
	GAME_ALREADY_ENDED
)

func NewApiError(e ErrorType) *ApiError {
	return e.getErrorByType()
}

func (et ErrorType) getErrorByType() *ApiError {
	switch et {
	case INVALID_USER_NAME_ERROR:
		message := "User name cannot be blank."
		statusCode := http.StatusBadRequest
		typee := "invalid_json"
		return &ApiError{Message: message, Type: typee, StatusCode: statusCode}
	case INVALID_PARAMETER_ERROR:
		message := "Invalid Parameter."
		statusCode := http.StatusBadRequest
		typee := "invalid_json"
		return &ApiError{Message: message, Type: typee, StatusCode: statusCode}
	case NO_RECORDS_FOUND_ERROR:
		message := "No Records Found."
		statusCode := http.StatusNotFound
		typee := "database"
		return &ApiError{Message: message, Type: typee, StatusCode: statusCode}
	case SAVING_GAME_ERROR:
		message := "Error when saving Game on database."
		statusCode := http.StatusInternalServerError
		typee := "database"
		return &ApiError{Message: message, Type: typee, StatusCode: statusCode}
	case SAVING_BOARD_ERROR:
		message := "Error when saving Board on database."
		statusCode := http.StatusInternalServerError
		typee := "database"
		return &ApiError{Message: message, Type: typee, StatusCode: statusCode}
	case INVALID_ACTION_ERROR:
		message := "Invalid Action. Valid actions are: \n flag \n click"
		statusCode := http.StatusBadRequest
		typee := "invalid_json"
		return &ApiError{Message: message, Type: typee, StatusCode: statusCode}
	case CUSTOM_ERROR:
		statusCode := http.StatusInternalServerError
		return &ApiError{StatusCode: statusCode}
	case CELL_ALREADY_CLICKED_ERROR:
		statusCode := http.StatusBadRequest
		message := "Cell Already Clicked!"
		typee := "invalid_json"
		return &ApiError{Message: message, Type: typee, StatusCode: statusCode}
	case LOST_GAME_ERROR:
		statusCode := http.StatusBadRequest
		message := "Game already Ended – you Lost!"
		typee := "invalid_json"
		return &ApiError{Message: message, Type: typee, StatusCode: statusCode}
	case CELL_ALREADY_FLAGGED_ERROR:
		statusCode := http.StatusBadRequest
		message := "Cell Already Flagged!"
		typee := "invalid_json"
		return &ApiError{Message: message, Type: typee, StatusCode: statusCode}
	case UPDATE_CELL_STATUS_ERROR:
		message := "Error Trying to update Cell Status."
		statusCode := http.StatusInternalServerError
		typee := "database"
		return &ApiError{Message: message, Type: typee, StatusCode: statusCode}
	case UPDATE_BOARD_ERROR:
		message := "Error Trying to update Board."
		statusCode := http.StatusInternalServerError
		typee := "database"
		return &ApiError{Message: message, Type: typee, StatusCode: statusCode}
	case UPDATE_GAME_ERROR:
		message := "Error Trying to update Game."
		statusCode := http.StatusInternalServerError
		typee := "database"
		return &ApiError{Message: message, Type: typee, StatusCode: statusCode}
	case GAME_ALREADY_ENDED:
		message := "Game Already Ended!"
		statusCode := http.StatusBadRequest
		typee := "invalid_json"
		return &ApiError{Message: message, Type: typee, StatusCode: statusCode}
	}

	return &ApiError{Message: "An unknown error occured.", StatusCode: http.StatusInternalServerError, Type: "server_error"}
}
