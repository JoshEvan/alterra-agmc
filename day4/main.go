package main

import (
	"github.com/JoshEvan/alterra-agmc-day4/config"
	"github.com/JoshEvan/alterra-agmc-day4/routes"

	"github.com/JoshEvan/alterra-agmc-day4/middlewares"
	echo "github.com/labstack/echo/v4"
)

func main() {
	initWebService()
}

func initWebService() {
	e := echo.New()
	router := routes.New(e, middlewares.InitAuth(e))
	middlewares.LogMiddleware(e)
	config.InitDB()
	router.Register()
	router.Start(":8080")
}
