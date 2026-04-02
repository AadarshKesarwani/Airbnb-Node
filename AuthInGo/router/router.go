package router

import (
    "github.com/go-chi/chi/v5"
	"AuthInGo/controllers"
    "github.com/go-chi/chi/v5/middleware"
    "AuthInGo/middlewares"
)

func SetupRouter(userRouter *UserRouter) *chi.Mux {

    chiRouter := chi.NewRouter()

    // chiRouter.Use(middlewares.RequestLogger) // Middleware for logging requests

	chiRouter.Use(middleware.Logger) // Built-in Chi middleware for
    //  logging requests
    chiRouter.Use(middlewares.RequestValidator) // Built-in Chi middleware for recovering from panics


    chiRouter.Get("/ping", controllers.PingHandler)


    userRouter.Register(chiRouter)

    return chiRouter
}