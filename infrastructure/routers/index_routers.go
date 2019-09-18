package routers

import (
	"github.com/K-jun1221/golang-server-template/interface/controllers"
	"github.com/labstack/echo"
)

func IndexRouting(e *echo.Echo) {
	e.GET("/", controllers.HelloEcho)
	e.GET("/csrf", controllers.CSRF)
	e.GET("/timeout", controllers.Timeout)
}
