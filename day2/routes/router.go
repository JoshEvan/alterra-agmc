package routes

import (
	"github.com/JoshEvan/alterra-agmc-day2/controllers"
	"github.com/labstack/echo/v4"
)

type Router struct {
	e *echo.Echo
}

func New(e *echo.Echo) Router {
	return Router{e: e}
}

func (r *Router) Start(port string) {
	r.e.Start(port)
}

func (r *Router) Register() {
	r.books()
	r.users()
}

func (r *Router) books() {
	r.e.GET("/books", controllers.GetBooks)
	r.e.GET("/book/:id", controllers.GetBookByID)
	r.e.POST("/books", controllers.CreateBook)
	r.e.PUT("/book/:id", controllers.UpdateBook)
	r.e.DELETE("/book/:id", controllers.DeleteBook)
}

func (r *Router) users() {
	r.e.GET("/users", controllers.GetUsers)
	r.e.GET("/user/:id", controllers.GetUserByID)
	r.e.POST("/users", controllers.CreateUser)
	r.e.PUT("/user/:id", controllers.UpdateUser)
	r.e.DELETE("/user/:id", controllers.DeleteUser)
}
