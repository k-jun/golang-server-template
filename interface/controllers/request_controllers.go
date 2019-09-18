package controllers

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

type body struct {
	Str            string  `json:"str"`
	Int            int64   `json:"int"`
	Float          float64 `json:"float"`
	WithUnderScore string  `json:"with_under_score"`
}

// {
// 	"str": "string body",
// 	"int": 1,
// 	"float": 1.1,
// 	"with_under_score": "be careful to underscore"
// }
func BindBody(c echo.Context) error {
	var params body
	err := c.Bind(&params)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, params)
}

type pathParams struct {
	Str   string  `json:"str"`
	Int   int64   `json:"int"`
	Float float64 `json:"float"`
}

// request/bind_path_params/:str/:int/:float
// request/bind_path_params/string/1/1.1
func BindPathParams(c echo.Context) error {
	s := c.Param("str")
	i := c.Param("int")
	f := c.Param("float")
	i64, err := strconv.ParseInt(i, 10, 64)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	f64, err := strconv.ParseFloat(f, 64)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	params := pathParams{Str: s, Int: i64, Float: f64}
	return c.JSON(http.StatusOK, params)
}

type queryParams struct {
	Str   string  `json:"str"`
	Int   int64   `json:"int"`
	Float float64 `json:"float"`
}

// request/bind_query_params
// request/bind_query_params?str=string_value&int=1&float=1.1
func BindQueryParams(c echo.Context) error {
	s := c.QueryParam("str")
	i := c.QueryParam("int")
	f := c.QueryParam("float")
	i64, err := strconv.ParseInt(i, 10, 64)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	f64, err := strconv.ParseFloat(f, 64)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	params := queryParams{Str: s, Int: i64, Float: f64}
	return c.JSON(http.StatusOK, params)
}
