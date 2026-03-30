package app

import (
    "AuthInGo/controllers"
    "AuthInGo/router"
    "AuthInGo/services"
    "AuthInGo/db/repositories"
    config "AuthInGo/config/env"
    "net/http"
    "fmt"
    "database/sql"
)


type Config struct {
    Addr string
}

func NewConfig() Config {
    port := config.GetString("PORT", ":8080")
    return Config{
        Addr: port,
    }
}

type Application struct {
    Config Config
    DB     *sql.DB
}

func NewApplication(cfg Config, db *sql.DB) *Application {
    return &Application{
        Config: cfg,
        DB:     db,
    }
}

func (app *Application) Run() error {
    ur := repositories.NewUserRepository()
    us := services.NewUserService(ur)
    uc := controllers.NewUserController(us)
    uRouter := router.NewUserRouter(uc)

    server := &http.Server{
        Addr:    app.Config.Addr,
        Handler: router.SetupRouter(uRouter),
    }
    fmt.Println("Starting server on", app.Config.Addr)
    return server.ListenAndServe()
}