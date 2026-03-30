package router

import (
    "AuthInGo/controllers"
    "github.com/go-chi/chi/v5"
)

type UserRouter struct {
    userController *controllers.UserController
}

func NewUserRouter(uc *controllers.UserController) *UserRouter {
    return &UserRouter{
        userController: uc,
    }
}

func (ur *UserRouter) Register(r chi.Router) {
    r.Post("/signup", ur.userController.RegisterUser)
}