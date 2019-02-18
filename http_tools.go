// Package tools implements helper methods to make development of rest APIs easy.
package tools

import (
	"encoding/json"
	"net/http"
)

// Unmarshal turns an http request body into the passed type.
func Unmarshal(v interface{}, r *http.Request) error {
	if r.Body == nil {
		return NoReqBodyErr
	}
	decoder := json.NewDecoder(r.Body)
	return decoder.Decode(&v)
}

// ServeJsonRes turns a struct into json format and writes the response along
// with an http status code.
func ServeJsonRes(w http.ResponseWriter, status int, res interface{}) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(res)
}
