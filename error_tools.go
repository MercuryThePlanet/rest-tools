package tools

/*
 * Adapted: https://blog.questionable.services/article/http-handler-error-handling-revisited/
 */

import (
	"encoding/json"
	"log"
	"net/http"
)

// Error represents a handler error. It provides methods for a HTTP status
// code and embeds the built-in error interface.
type Error interface {
	error
	Status() int
}

// StatusError represents an error with an associated HTTP status code.
type StatusError struct {
	Code int
	Err  error
}

// Allows StatusError to satisfy the error interface.
func (se StatusError) Error() string {
	return se.Err.Error()
}

// Returns HTTP status code.
func (se StatusError) Status() int {
	return se.Code
}

// ErrorHandle is a modified http.Handle which returns an error.
type ErrorHandle = func(w http.ResponseWriter, r *http.Request) error

// ErrorHandler Stores an ErrorHandle to be used during ServeHTTP
type ErrorHandler struct {
	H func(w http.ResponseWriter, r *http.Request) error
}

// ServeHTTP allows our Handler type to satisfy http.Handler.
func (eh ErrorHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if err := eh.H(w, r); err != nil {
		w.Header().Set("Content-Type", "application/json")
		var errRes struct {
			Error string `json:"error"`
		}
		switch e := err.(type) {
		case Error:
			log.Printf("HTTP %d - %s", e.Status(), e)
			w.WriteHeader(e.Status())
			errRes.Error = e.Error()
		default:
			w.WriteHeader(http.StatusInternalServerError)
			errRes.Error = http.StatusText(http.StatusInternalServerError)
		}
		json.NewEncoder(w).Encode(errRes)
	}
}
