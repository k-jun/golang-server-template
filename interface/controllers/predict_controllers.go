package controllers

import (
	"net/http"

	"github.com/labstack/echo"
)

func PredictIris(c echo.Context) error {
	return c.String(http.StatusOK, "this api is working! Ver 1.0.0")
}
