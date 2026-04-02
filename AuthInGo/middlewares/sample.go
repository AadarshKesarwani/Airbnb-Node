package middlewares

import (
	"fmt"
	"net/http"
)
//to define a middleware in Go, we create a function that takes an http.Handler as an argument and returns a new http.Handler. The middleware can perform any necessary processing before or after calling the next handler in the chain. In this example, we are creating a simple logging middleware that prints the HTTP method and URL of each incoming request.

func RequestLogger(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		fmt.Println("Received request:", r.Method, r.URL.Path)
		next.ServeHTTP(w, r) // Call the next handler in the chain
	})
}