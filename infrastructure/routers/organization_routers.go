package routers

import (
	"github.com/K-jun1221/golang-server-template/interface/controllers"
	"github.com/labstack/echo"
)

func OrganizationRouting(e *echo.Group) {
	e.GET("/", controllers.GetOrganizations)
	e.POST("/", controllers.CreateOrganization)
	e.PUT("/:id", controllers.UpdateOrganization)
	e.DELETE("/:id", controllers.DeleteOrganization)
}
