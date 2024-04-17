package utils

import "net/http"

// RespondWithError sends an HTTP error response with the provided status code and message
func RespondWithError(w http.ResponseWriter, statusCode int, message string) {
	w.WriteHeader(statusCode)
	w.Write([]byte(message))
}
