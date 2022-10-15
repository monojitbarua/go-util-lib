package error

import "net/http"

type ApplicationError struct {
	Code    int    `json:",omitempty"`
	Message string `json:"message"`
}

func (e ApplicationError) AsMessage() *ApplicationError {
	return &ApplicationError{Message: e.Message}
}

func NotFoundError(message string) *ApplicationError {
	return &ApplicationError{Message: message, Code: http.StatusNotFound}
}

func UnexpectedError(message string) *ApplicationError {
	return &ApplicationError{Message: message, Code: http.StatusInternalServerError}
}

func ValidationError(message string) *ApplicationError {
	return &ApplicationError{Message: message, Code: http.StatusUnprocessableEntity}
}

func AuthenticationError(message string) *ApplicationError {
	return &ApplicationError{Message: message, Code: http.StatusUnauthorized}
}

func AuthorizationError(message string) *ApplicationError {
	return &ApplicationError{Message: message, Code: http.StatusForbidden}
}
