package main

import (
	"github.com/JoshEvan/alterra-agmc-day2/config"
	"github.com/JoshEvan/alterra-agmc-day2/routes"
	echo "github.com/labstack/echo/v4"
)

func main() {
	initWebService()
}

func initWebService() {
	router := routes.New(echo.New())
	config.InitDB()
	router.Register()
	router.Start(":8080")
}
