package router

import (
    "github.com/go-chi/chi/v5"
	"AuthInGo/controllers"
)

func SetupRouter(userRouter *UserRouter) *chi.Mux {
    chiRouter := chi.NewRouter()
    chiRouter.Get("/ping", controllers.PingHandler)
    userRouter.Register(chiRouter)
    return chiRouter
}