package tools

import (
	"net/http"
)

// ErrorHandle is a modified http.Handle which returns an error.
type ErrorHandle = func(w http.ResponseWriter, r *http.Request) error

// MethodMap maps an http.Method string to an ErrorHandle.
type MethodMap = map[string]ErrorHandle

// RestHelper is a helper struct to make rest api method handling easier.
type RestHelper struct {
	methods MethodMap
}

// Returns a new instance of RestHelper given a MethodMap.
func NewRestHelper(methods MethodMap) *RestHelper {
	return &RestHelper{methods}
}

// Adds a method to RestHelper's MethodMap.
func (rh *RestHelper) AddMethod(m string, h ErrorHandle) {
	rh.methods[m] = h
}

// Adds all methods in the passed MethodMap to RestHelper's MethodMap.
func (rh *RestHelper) AddMethods(methods MethodMap) {
	for k, v := range methods {
		rh.methods[k] = v
	}
}

// Handler will run the appropriate handler for any given method. If the
// method has not been set the handler returns a method not allowed error.
func (rh *RestHelper) Handler(w http.ResponseWriter, r *http.Request) error {
	handle, ok := rh.methods[r.Method]
	if !ok {
		return MethodNotAllowedErr
	}
	return handle(w, r)
}
