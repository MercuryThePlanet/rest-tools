package tools

import (
	"errors"
)

var (
	// http errors
	MethodNotAllowedErr = errors.New("Method not allowed.")
	NoReqBodyErr        = errors.New("No request body.")

	// func http errors
	ParamRequiredErr = func(k string) error {
		return errors.New("Query Parameter `" + k + "` is required.")
	}
	ParamWrongTypeErr = func(t string) error {
		return errors.New("Parameter must be a " + t + ".")
	}
)
