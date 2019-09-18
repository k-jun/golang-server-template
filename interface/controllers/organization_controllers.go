package controllers

import (
	"net/http"
	"strconv"

	"github.com/K-jun1221/golang-server-template/infrastructure/datastore/mysql"
	"github.com/K-jun1221/golang-server-template/infrastructure/datastore/mysql/models"
	"github.com/labstack/echo"
	"github.com/volatiletech/sqlboiler/boil"
)

type editableColumn struct {
	Name string `json:"name"`
}

func GetOrganizations(c echo.Context) error {
	orgs, err := models.Organizations().All(mysql.CTX, mysql.DB)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, orgs)
}

func CreateOrganization(c echo.Context) error {
	var params editableColumn
	err := c.Bind(&params)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	newOrg := models.Organization{
		Name: params.Name,
	}

	err = newOrg.Insert(mysql.CTX, mysql.DB, boil.Infer())
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, newOrg)
}

func UpdateOrganization(c echo.Context) error {
	id := c.Param("id")
	var params editableColumn

	idint, err := strconv.Atoi(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	err = c.Bind(&params)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	org, err := models.FindOrganization(mysql.CTX, mysql.DB, idint)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	org.Name = params.Name

	_, err = org.Update(mysql.CTX, mysql.DB, boil.Infer())
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, org)
}

func DeleteOrganization(c echo.Context) error {
	id := c.Param("id")

	idint, err := strconv.Atoi(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	org, err := models.FindOrganization(mysql.CTX, mysql.DB, idint)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	result, err := org.Delete(mysql.CTX, mysql.DB)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, result)
}
