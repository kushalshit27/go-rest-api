package utils

import (
	"net/http"
)

// Route route
type Route struct {
	Path    string
	Method  string
	Handler http.HandlerFunc
}

// Routes routes
type Routes []Route

// Error error
type Error struct {
	Error   bool   `json:"error"`
	Message string `json:"message"`
}

// Response as
type Response struct {
	Error Error       `json:"error"`
	Data  interface{} `json:"data"`
}

// Data data response
type Data struct {
	Data interface{} `json:"data"`
}

// AddRoute AddRoute
func AddRoute(path string, method string, handler http.HandlerFunc) Route {
	return Route{path, method, handler}
}

// ResponseSuccess ResponseSuccess
func ResponseSuccess(message string, i interface{}) Response {
	return Response{Error{false, message}, i}
}

// ResponseError ResponseError
func ResponseError(message string) Response {
	return Response{Error{true, message}, nil}
}
