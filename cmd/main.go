package main

import (
	"cinema-ticket/config"
	"cinema-ticket/internal/api/server"
	postgres "cinema-ticket/internal/repostiry"
)


func main() {

cfg , err := config.Load("./config/config.yaml")
if err != nil {
	panic(err)
}
postgres.InitDB(cfg.Database)
postgres.AutoMigrate()
server.Start(cfg.Server)


}
