package middlewares

import (
	env "AuthInGo/config/env"
	"context"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/golang-jwt/jwt/v5"
)

//explain each steps

// JWTAuthMiddleware is a middleware that checks for the presence of a valid JWT token in the Authorization header of incoming requests. It performs the following steps:

// 1. It retrieves the Authorization header from the incoming request. If the header is missing, it returns a 401 Unauthorized response.

// 2. It checks if the Authorization header starts with the "Bearer " prefix. If it does not, it returns a 401 Unauthorized response.

// 3. It extracts the token string by removing the "Bearer " prefix from the Authorization header. If the token string is empty, it returns a 401 Unauthorized response.

// 4. It initializes a jwt.MapClaims object to hold the claims from the token.

//claims are the pieces of information encoded in the JWT token, such as user ID and email.

// 5. It parses the token string using the jwt.ParseWithClaims function, passing in the token string, the claims object, and a function that returns the secret key used to sign the token. If there is an error during parsing (e.g., invalid token), it returns a 401 Unauthorized response.

// 6. It retrieves the user ID and email from the claims. If either of these values is missing or of the wrong type, it returns a 401 Unauthorized response.

// 7. It prints the user ID and email to the console for debugging purposes.

// 8. It creates a new context with the user ID and email values and calls the next handler in the chain, passing the modified request with the new context.

// In summary, this middleware ensures that incoming requests contain a valid JWT token and extracts the user information from the token to be used in subsequent handlers. If the token is missing or invalid, it prevents access to protected routes by returning an appropriate error response.




func JWTAuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){

		// Retrieve the Authorization header from the incoming request
		authHeader := r.Header.Get("Authorization")

		// Check if the Authorization header is present

		// If the header is missing, return a 401 Unauthorized response
		if authHeader == "" {
			http.Error(w, "Missing authorization header", http.StatusUnauthorized)
			return
		}

		// Check if the Authorization header starts with the "Bearer " prefix

		// If it does not, return a 401 Unauthorized response
		if !strings.HasPrefix(authHeader, "Bearer ") {
			http.Error(w, "Invalid authorization header", http.StatusUnauthorized)
			return
		}

		// Extract the token string by removing the "Bearer " prefix from the Authorization header

		

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		// If the token string is empty, return a 401 Unauthorized response
		if tokenString == "" {
			http.Error(w, "Invalid authorization header", http.StatusUnauthorized)
			return
		}

		// Initialize a jwt.MapClaims object to hold the claims from the token

		//claims = the pieces of information encoded in the JWT token, such as user ID and email.
		//example of claims: {"user_id": 123, "email": "user@example.com"}
		claims := jwt.MapClaims{}

		// Parse the token string using the jwt.ParseWithClaims function, passing in the token string, the claims object, and a function that returns the secret key used to sign the token. If there is an error during parsing (e.g., invalid token), return a 401 Unauthorized response.
		_, err := jwt.ParseWithClaims(tokenString, &claims, func(token *jwt.Token) (interface{}, error) {
			return []byte(env.GetString("JWT_SECRET_KEY", "TOKEN")), nil
		})


		// If there is an error during parsing (e.g., invalid token), return a 401 Unauthorized response.


		if err != nil {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		// Retrieve the user ID and email from the claims. If either of these values is missing or of the wrong type, return a 401 Unauthorized response.
		userId, okId := claims["user_id"].(float64)

		email, okEmail := claims["email"].(string)


		// If either of these values is missing or of the wrong type, return a 401 Unauthorized response.
		if !okId || !okEmail {
			http.Error(w, "Invalid token claims", http.StatusUnauthorized)
			return
		}

		// Print the user ID and email to the console for debugging purposes.

		fmt.Println("User ID from token:", userId)
		fmt.Println("Email from token:", email)

		// Create a new context with the user ID and email values and call the next handler in the chain, passing the modified request with the new context.

		ctx := context.WithValue(r.Context(), "user_id", strconv.FormatFloat(userId, 'f', -1, 64))

		// strconv.FormatFloat is used to convert the userId from float64 to string format, which is more suitable for storing in the context and later retrieving it in handlers.

		// The user ID is stored in the context with the key "user_id", and the email is stored with the key "email". This allows subsequent handlers to access these values from the context when processing the request.

		
		ctx = context.WithValue(ctx, "email", email)


		// By calling next.ServeHTTP with the modified request (r.WithContext(ctx)), the middleware ensures that the user ID and email are available to any handlers that are executed after this middleware in the request processing chain.
		next.ServeHTTP(w, r.WithContext(ctx))

	})
}