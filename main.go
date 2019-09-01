package main

import (
	"github.com/K-jun1221/golang-server-template/infrastructure/routers"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))

	routers.IndexRouting(e)
	routers.PredictRouting(e.Group("/predict"))

	e.Logger.Fatal(e.Start(":8080"))
}
