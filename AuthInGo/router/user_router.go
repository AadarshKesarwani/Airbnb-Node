package router

import (
    "AuthInGo/controllers"
    "AuthInGo/middlewares"
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


    r.With(middlewares.JWTAuthMiddleware).Get("/users/{id}", ur.userController.GetUserById)

    r.With(middlewares.UserCreateRequestValidator).Post("/signup", ur.userController.CreateUser)

    
    r.With(middlewares.UserLoginRequestValidator).Post("/login", ur.userController.LoginUser)
}