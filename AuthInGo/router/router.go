package router

import (
    "AuthInGo/controllers"
    // "AuthInGo/middlewares"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)


type Router interface {
	Register(r chi.Router)
}



func SetupRouter(userRouter *UserRouter) *chi.Mux {

    chiRouter := chi.NewRouter()

    // chiRouter.Use(middlewares.RequestLogger) // Middleware for logging requests

	chiRouter.Use(middleware.Logger) // Built-in middleware for logging

    // chiRouter.Use(middlewares.RateLimiteMiddleware) // Middleware for rate limiting

    chiRouter.Get("/ping", controllers.PingHandler)


    userRouter.Register(chiRouter)

    return chiRouter
}