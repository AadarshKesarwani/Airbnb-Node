package utils

import (
    "encoding/json"
    "net/http"

    v "github.com/go-playground/validator/v10"
)


// JSONResponse is a helper function to send JSON responses

//first we create variable validate of type *validator.Validate, which is a pointer to the validator struct from the go-playground/validator package. This variable will be used to perform validation on our data structures.


var validate *v.Validate

//init function is a special function in Go that is automatically executed when the package is initialized. In this case, we are initializing the validator variable by creating a new instance of the validator using validator.New(). This allows us to use the validator variable throughout our package to perform validation on our data structures.
func init() {
	validate = NewValidator()
}


//new function NewValidator is a helper function that creates and returns a new instance of the validator with the option WithRequiredStructEnabled. This option enables the validation of struct fields that are marked as required. By calling this function, we can get a new validator instance that is configured to validate required struct fields, which can be useful for ensuring that our data structures meet certain validation criteria before processing them further.


func NewValidator() *v.Validate  {
	return v.New(v.WithRequiredStructEnabled())
}


//ValidateStruct is a helper function that validates the provided struct using the validator instance. It takes a payload of type any (which can be any struct) and returns an error if the validation fails. This function can be used to validate request payloads in our handlers before processing them further.

func ValidateStruct(payload any) error {
	return validate.Struct(payload)
}


func WriteJSONResponse(w http.ResponseWriter, statusCode int, data any) error {

	// Set the Content-Type header to application/json
	w.Header().Set("Content-Type", "application/json")

	// Set the status code for the response

	w.WriteHeader(statusCode)

	// Encode the data as JSON and write it to the response

	//NewEncoder is a function from the encoding/json package that creates a new JSON encoder that writes to the provided io.Writer, which in this case is the http.ResponseWriter (w). 
	

	// The Encode method of the encoder is then called with the data to be encoded as JSON. This will convert the data into JSON format and write it to the response body. If there is an error during encoding, it will be returned to the caller for handling.
	return json.NewEncoder(w).Encode(data)
}


func WriteSuccessResponse(w http.ResponseWriter, data any, status int, message string) error {
	response := map[string]any{}
	response["status"] = "success"
	response["data"] = data
	response["message"] = message

	return WriteJSONResponse(w, status, response)
}


func WriteErrorResponse(w http.ResponseWriter, status int, message string, err error) error {
	response := map[string]any{}
	response["status"] = "error"
	response["message"] = message
	response["error"] = err.Error()
	return WriteJSONResponse(w, status, response)
}



//function to read json body

func ReadJSONBody(r *http.Request, result any) error {

	// Decode the JSON body into the provided result variable

	decoder := json.NewDecoder(r.Body)// Create a new JSON decoder that reads from the request body (r.Body). This allows us to read the incoming JSON data from the request.

	decoder.DisallowUnknownFields() // This line tells the decoder to disallow any unknown fields in the JSON data. If the incoming JSON contains fields that are not defined in the struct we are decoding into, it will return an error. This is a good practice to ensure that we only accept expected fields and avoid potential issues with unexpected data.

	// The Decode method of the decoder is then called with the result variable, which is where the decoded JSON data will be stored. If there is an error during decoding, it will be returned to the caller for handling.

	return decoder.Decode(result)
}