package controllers

import (
    "net/http"
    "AuthInGo/services"
    "net/http"
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
	fmt.Println("GetUserById called in UserController")
	uc.UserService.GetUserById()
	w.Write([]byte("User fetching endpoint done"))
}