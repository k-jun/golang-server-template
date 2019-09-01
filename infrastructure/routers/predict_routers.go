package routers

import (
	"github.com/K-jun1221/golang-server-template/interface/controllers"
	"github.com/labstack/echo"
)

func PredictRouting(e *echo.Group) {
	e.GET("", controllers.PredictIris)
}
