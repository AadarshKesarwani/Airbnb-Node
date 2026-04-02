package controllers

import (
    "AuthInGo/dto"
    "AuthInGo/utils"
    "AuthInGo/services"
    "net/http"
    "fmt"
    "strconv"
)




type UserController struct {
    UserService services.UserService
}




func NewUserController(us services.UserService) *UserController {
    return &UserController{
        UserService: us,
    }
}





func (uc *UserController) GetUserById(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Fetching user by ID in UserController")
	// extract userid from url parameters

    // In this example, we are using the standard library's net/http package to handle HTTP requests. We extract the user ID from the URL query parameters, validate it, and then call the UserService to fetch the user details. We also handle various error scenarios and return appropriate HTTP responses.


	userId := r.URL.Query().Get("id")

    // Validate the user ID
    // We check if the user ID is provided and if it can be parsed into an integer. If the user ID is missing or invalid, we return a 400 Bad Request response with an appropriate error message. If the user is not found, we return a 404 Not Found response. If there is an internal server error while fetching the user, we return a 500 Internal Server Error response. If the user is successfully fetched, we return a 200 OK response with the user details in the response body.
	if userId == "" {
		utils.WriteErrorResponse(w, http.StatusBadRequest, "User ID is required", fmt.Errorf("missing user ID"))
		return
	}

    // Parse the user ID to an integer
	id, err := strconv.ParseInt(userId, 10, 64)

    // If there was an error parsing the user ID, return a 400 Bad Request response
	if err != nil {
		utils.WriteErrorResponse(w, http.StatusBadRequest, "Invalid user ID format", err)
		return
	}

    // Call the UserService to fetch the user details by ID
	user, err := uc.UserService.GetUserById(id)

    // Handle errors and return appropriate responses
	if err != nil {
		utils.WriteErrorResponse(w, http.StatusInternalServerError, "Failed to fetch user", err)
		return
	}

    // If the user is not found, return a 404 Not Found response
	if user == nil {
		utils.WriteErrorResponse(w, http.StatusNotFound, "User not found", fmt.Errorf("user with ID %d not found", id))
		return
	}

    // If the user is successfully fetched, return a 200 OK response with the user details
	utils.WriteSuccessResponse(w, user, http.StatusOK, "User fetched successfully")
	fmt.Println("User fetched successfully:", user)
}





func (uc *UserController) CreateUser(w http.ResponseWriter, r *http.Request) {

    fmt.Println("CreateUser called in UserController")
    // Here we would typically parse the request payload, validate it, and then call the UserService to perform the user creation logic. For simplicity, we are just printing a message and returning a response.

    //payload is the data that we receive in the request body when a client makes a request to create a new user. We need to read this data and parse it into a structured format (like a struct) so that we can work with it in our code. In this example, we are using the utils.ReadJSONBody function to read the JSON payload from the request body and unmarshal it into a CreateUserRequestDTO struct. This allows us to easily access the user details (like username, email, and password) that were sent in the request when we call the UserService to create the user.
    var payload dto.CreateUserRequestDTO

    // Read and parse the JSON payload from the request body
    // We call the utils.ReadJSONBody function, passing in the HTTP request (r) and a pointer to the payload variable. This function will read the JSON data from the request body and unmarshal it into the CreateUserRequestDTO struct. If there is an error during this process (e.g., if the JSON is malformed), we return a 400 Bad Request response with an appropriate error message.
    if jsonErr := utils.ReadJSONBody(r, &payload); jsonErr != nil {
        utils.WriteErrorResponse(w, http.StatusBadRequest, "something went wrong", jsonErr)
        return
    }

    // Print the parsed payload for debugging purposes
    fmt.Println("Parsed create user payload:", payload)


    user,err := uc.UserService.CreateUser(&payload)
    if err != nil {
        utils.WriteErrorResponse(w, http.StatusInternalServerError, "user creation failed", err)
        return
    }

    utils.WriteSuccessResponse(w, user, http.StatusCreated, "User created successfully")
}






func (uc *UserController) LoginUser(w http.ResponseWriter, r *http.Request) {

    fmt.Println("LoginUser called in UserController")

    // Here we would typically parse the login request payload, validate it, and then call the UserService to perform the login logic. For simplicity, we are just printing a message and returning a response.
    var payload dto.LoginUserRequestDTO

    if jsonErr := utils.ReadJSONBody(r, &payload); jsonErr != nil {
        utils.WriteErrorResponse(w, http.StatusBadRequest, "something went wrong", jsonErr)
        return  
    }

    fmt.Println("Parsed login payload:", payload)


    // Call the UserService to perform the login logic
    token, err := uc.UserService.LoginUser(&payload)

    if err != nil {
        utils.WriteErrorResponse(w, http.StatusInternalServerError, "login failed", err)
        return
    }

    // If login is successful, return the token in the response
    utils.WriteSuccessResponse(w, token, http.StatusOK, "user logged in successfully")

}