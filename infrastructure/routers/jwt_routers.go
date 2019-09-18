package routers

import (
	"github.com/K-jun1221/golang-server-template/interface/controllers"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func JwtRouting(e *echo.Group) {
	// jwtToken対策
	jwtMiddleware := middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey:  []byte("hwatXTgp1L"),
		ContextKey:  "current_user",
		TokenLookup: "header:X-ECHO-JWT-TOKEN",
	})
	e.GET("/login", controllers.Login)
	e.GET("/mypage", controllers.MyPage, jwtMiddleware)
	e.GET("/logout", controllers.Logout, jwtMiddleware)
}
