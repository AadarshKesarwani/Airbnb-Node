package controllers

import (
    "net/http"
    "AuthInGo/services"
)

type UserController struct {
    UserService services.UserService
}

func NewUserController(us services.UserService) *UserController {
    return &UserController{
        UserService: us,
    }
}

func (uc *UserController) RegisterUser(w http.ResponseWriter, r *http.Request) {
    uc.UserService.CreateUser()
    w.Write([]byte("User registration endpoint"))
}