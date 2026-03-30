package main


import (
    "AuthInGo/app"
    dbConfig "AuthInGo/config/db"
    config "AuthInGo/config/env"
    "log"
)




func main() {
    config.Load()
    cfg := app.NewConfig()
    
    db, err := dbConfig.SetupDB()
    if err != nil {
        log.Fatal(err)
    }
    
    application := app.NewApplication(cfg, db)
    application.Run()
}