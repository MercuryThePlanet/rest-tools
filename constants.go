package tools

import (
	"errors"
)

var (
	// http errors
	MethodNotAllowedErr = errors.New("Method not allowed.")
	NoReqBodyErr        = errors.New("No request body.")
)
