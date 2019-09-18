package routers

import (
	"github.com/K-jun1221/golang-server-template/interface/controllers"
	"github.com/labstack/echo"
)

func RequestRouting(e *echo.Group) {
	e.POST("/bind_body", controllers.BindBody)
	e.GET("/bind_query_params", controllers.BindQueryParams)
	e.GET("/bind_path_params/:str/:int/:float", controllers.BindPathParams)
}
