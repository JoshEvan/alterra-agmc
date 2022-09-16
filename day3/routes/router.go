package routes

import (
	"github.com/JoshEvan/alterra-agmc-day3/controllers"
	"github.com/labstack/echo/v4"
)

type Router struct {
	e         *echo.Echo
	authGroup *echo.Group
}

func New(e *echo.Echo, authGroup *echo.Group) Router {
	return Router{e: e, authGroup: authGroup}
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
	r.e.POST("/login", controllers.LoginUser)
	r.e.POST("/users", controllers.CreateUser)

	r.authGroup.GET("/users", controllers.GetUsers)
	r.authGroup.GET("/user/:id", controllers.GetUserByID)
	r.authGroup.PUT("/user/:id", controllers.UpdateUser)
	r.authGroup.DELETE("/user/:id", controllers.DeleteUser)
}
