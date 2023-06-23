package controllers

type HttpError struct {
	Message string `json:"message"`
}

func newHttpError(message string) *HttpError {
	return &HttpError{
		Message: message,
	}
}
