package app

import (
    dbConfig "AuthInGo/config/db"
    "AuthInGo/controllers"
    "AuthInGo/router"
    "AuthInGo/services"
    repo "AuthInGo/db/repositories"
    config "AuthInGo/config/env"
    "fmt"
    "database/sql"
    "time"
    "net/http"
)



type Config struct {
    Addr string //PORT
}



func NewConfig() Config {
    port := config.GetString("PORT", ":8080")
    return Config{
        Addr: port,
    }
}

type Application struct {
    Config Config
}

func NewApplication(cfg Config) *Application {
    return &Application{
        Config: cfg,
    }
}

func (app *Application) Run() error {
    db , err := dbConfig.SetupDB()
    if err != nil {
        fmt.Println("Failed to set up database:", err)
        return err
    }
    ur := repo.NewUserRepository()
    us := services.NewUserService(ur)
    uc := controllers.NewUserController(us)
    uRouter := router.NewUserRouter(uc)

    server := &http.Server{
        Addr:    app.Config.Addr,
        Handler: router.SetupRouter(uRouter),
        ReadTimeout:  10 * time.Second,
        WriteTimeout: 10 * time.Second,
    }
    fmt.Println("Starting server on", app.Config.Addr)
    return server.ListenAndServe()
}