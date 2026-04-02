package middlewares

import (
	"AuthInGo/dto"
	"AuthInGo/utils"
	"context"
	"fmt"
	"net/http"
)


func UserLoginRequestValidator(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var payload dto.LoginUserRequestDTO // Define the type of payload you expect

		// Read and decode the JSON body into the payload
		if err := utils.ReadJSONBody(r, &payload); err != nil {
			utils.WriteErrorResponse(w, http.StatusBadRequest, "Invalid request body", err)
			return
		}

		// Validate the payload using the Validator instance
		if err := utils.ValidateStruct(payload); err != nil {
			utils.WriteErrorResponse(w, http.StatusBadRequest, "Validation failed", err)
			return
		}

		fmt.Println("UserLoginRequestValidator passed validation for payload:", payload)

		// Store the validated payload in the request context for later use in the handler
		ctx := context.WithValue(r.Context(), "payload", payload)
		r = r.WithContext(ctx)

		next.ServeHTTP(w, r) // Call the next handler in the chain
	})
}




func UserCreateRequestValidator(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var payload dto.CreateUserRequestDTO // Define the type of payload you expect	

		// Read and decode the JSON body into the payload
		if err := utils.ReadJSONBody(r, &payload); err != nil {
			utils.WriteErrorResponse(w, http.StatusBadRequest, "Invalid request body", err)
			return
		}

		// Validate the payload using the Validator instance
		if err := utils.ValidateStruct(payload); err != nil {
			utils.WriteErrorResponse(w, http.StatusBadRequest, "Validation failed", err)
			return
		}

		fmt.Println("UserCreateRequestValidator passed validation for payload:", payload)
		// Store the validated payload in the request context for later use in the handler
		ctx := context.WithValue(r.Context(), "payload", payload)
		r = r.WithContext(ctx)	

		next.ServeHTTP(w, r) // Call the next handler in the chain
	})
}


