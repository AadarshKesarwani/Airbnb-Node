package router

import (
    "AuthInGo/controllers"
    "github.com/go-chi/chi/v5"
)

type UserRouter struct {
    userController *controllers.UserController
}

func NewUserRouter(_userController *controllers.UserController) *UserRouter {
    return &UserRouter{
        userController: _userController,
    }
}

func (ur *UserRouter) Register(r chi.Router) {

    r.Get("/users/{id}", ur.userController.GetUserById)

    r.Post("/signup", ur.userController.CreateUser)
  
    r.Post("/login", ur.userController.LoginUser)
}