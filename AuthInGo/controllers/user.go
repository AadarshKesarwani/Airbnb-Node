package controllers

import (
    "AuthInGo/dto"
    "AuthInGo/utils"
    "AuthInGo/services"
    "net/http"
    "fmt"
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
		userId =r.Context().Value("userID").(string)
	}

    fmt.Println("User ID extracted from URL:", userId)

    if userId == "" {
        utils.WriteErrorResponse(w, http.StatusBadRequest, "User ID is required", fmt.Errorf("user ID is missing in the request"))
        return
    }

    user, err := uc.UserService.GetUserById(userId)

    if err != nil {
        utils.WriteErrorResponse(w, http.StatusInternalServerError, "Failed to fetch user", err)
        return
    }

    // If the user is not found, return a 404 Not Found response
	if user == nil {
		utils.WriteErrorResponse(w, http.StatusNotFound, "User not found", fmt.Errorf("user with ID %s not found", userId))
		return
	}

    // If the user is successfully fetched, return a 200 OK response with the user details
	utils.WriteSuccessResponse(w, user, http.StatusOK, "User fetched successfully")
	fmt.Println("User fetched successfully:", user)
}





func (uc *UserController) CreateUser(w http.ResponseWriter, r *http.Request) {

    fmt.Println("CreateUser called in UserController")
    
    payload := r.Context().Value("payload").(dto.CreateUserRequestDTO)

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

    payload := r.Context().Value("payload").(dto.LoginUserRequestDTO)

    // Print the parsed payload for debugging purposes
    fmt.Println("Parsed login user payload:", payload)


    // Call the UserService to perform the login logic
    jwtToken, err := uc.UserService.LoginUser(&payload)

    if err != nil {
        utils.WriteErrorResponse(w, http.StatusInternalServerError, "login failed", err)
        return
    }

    // If login is successful, return the token in the response
    utils.WriteSuccessResponse(w, jwtToken, http.StatusOK, "user logged in successfully")

}