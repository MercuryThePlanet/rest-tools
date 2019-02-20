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

// PathParamToInt takes the urlPath and section section to trim, then converts
// the remainder into an int
func PathParamToInt(urlPath, trim) (res int, err error) {
	v := strings.TrimPrefix(urlPath, trim)
	res, err := strconv.Atoi(v)
	if err != nil {
		return 0, errors.New("URL Path Parameter must be an integer.")
	}
	return
}

// QueryParamToFloat64 takes a query param key and converts the value to float64
func QueryParamToFloat64(k string, r *http.Request) (res float64, err error) {
	v, ok := r.URL.Query()[k]
	if !ok {
		return 0, errors.New("Query Parameter `" + k + "` is required.")
	}
	res, err := strconv.ParseFloat(v[0], 64)
	if err != nil {
		return 0, errors.New("Query Parameter `" + k + "` must be a float64.")
	}
	return
}
