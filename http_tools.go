// Package tools implements helper methods to make development of rest APIs easy.
package tools

import (
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
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
func PathParamToInt(urlPath, trim string) (res int, err error) {
	v := strings.TrimPrefix(urlPath, trim)
	res, err = strconv.Atoi(v)
	if err != nil {
		return 0, ParamRequiredErr("integer")
	}
	return
}

// QueryParamToFloat64 takes a query param key and converts the value to float64
func QueryParamToFloat64(k string, r *http.Request) (res float64, err error) {
	v, ok := r.URL.Query()[k]
	if !ok {
		return 0, ParamRequiredErr(k)
	}
	res, err = strconv.ParseFloat(v[0], 64)
	if err != nil {
		return 0, ParamWrongTypeErr("float64")
	}
	return
}

// QueryParamToString takes a query param key and gets the string value
func QueryParamToString(k string, r *http.Request) (res string, err error) {
	a, ok := r.URL.Query()[k]
	if !ok {
		return "", ParamRequiredErr(k)
	}
	res = a[0]
	return
}

// GenerateToken creates a url-safe token of n-bytes
func GenerateToken(n int) (string, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(b), nil
}
