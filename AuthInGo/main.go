package main


import (
    "AuthInGo/app"
    config "AuthInGo/config/env"
    "log"
)




func main() {
    config.Load()
    cfg := app.NewConfig()
    
    application := app.NewApplication(cfg)
    application.Run()
}