package middlewares

import (
	"AuthInGo/utils"
	"net/http"
)

func RequestValidator(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var payload any // Define the type of payload you expect

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

		next.ServeHTTP(w, r) // Call the next handler in the chain
	})
}