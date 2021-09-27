package main

import (
	"go-ecommerce/config"
	"go-ecommerce/routes"
)


func main() {
	config.InitDB()
	ServerHost := config.GetEnvVariable("SERVER_HOST")
	e := routes.New()
	e.Logger.Fatal(e.Start(":"+ServerHost))
}
