package routers

import (
	"github.com/K-jun1221/golang-server-template/interface/controllers"
	"github.com/labstack/echo"
)

func CookieRouting(e *echo.Group) {
	e.GET("/set_cookie", controllers.SetCookie)
	e.GET("/get_cookie", controllers.GetCookie)
	e.GET("/get_cookies", controllers.GetCookies)
}
