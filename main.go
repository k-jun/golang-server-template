package main

import (
	"github.com/K-jun1221/golang-server-template/infrastructure/datastore/mysql"
	"github.com/K-jun1221/golang-server-template/infrastructure/routers"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	e := echo.New()
	err := mysql.Initialize()
	if err != nil {
		panic(err.Error())
	}

	// CORS
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))

	// BasicAuth
	// e.Use(middleware.BasicAuth(func(username, password string, c echo.Context) (bool, error) {
	// 	if username == "root" && password == "pass" {
	// 		return true, nil
	// 	}
	// 	return false, nil
	// }))

	// CSRF
	e.Use(middleware.CSRFWithConfig(middleware.CSRFConfig{
		TokenLookup: "header:X-ECHO-CSRF-TOKEN",
	}))

	// Logger
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))

	// Recover
	// e.Use(middleware.RecoverWithConfig(middleware.RecoverConfig{
	// 	StackSize: 1 << 10, // 1 KB
	// }))

	routers.IndexRouting(e)
	routers.CookieRouting(e.Group("/cookie"))
	routers.RequestRouting(e.Group("/request"))
	routers.JwtRouting(e.Group("/jwt"))
	routers.OrganizationRouting(e.Group("/organization"))
	e.Start(":8080")
}
