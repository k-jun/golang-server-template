package controllers

import (
	"net/http"
	"time"

	"github.com/labstack/echo"
)

const cookieValue = "golang-server-template"
const cookieName = "applicationName"

// TODO Comment
func SetCookie(c echo.Context) error {
	cookie, err := c.Cookie(cookieName)
	if err != nil {
		cookie := new(http.Cookie)
		cookie.Name = cookieName
		cookie.Value = cookieValue
		cookie.Expires = time.Now().Add(24 * time.Hour)
		c.SetCookie(cookie)
		return c.JSON(http.StatusOK, cookie)
	}
	return c.JSON(http.StatusOK, cookie)
}

// TODO Comment
func GetCookie(c echo.Context) error {
	cookie, err := c.Cookie(cookieName)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, cookie)
}

// TODO Comment
func GetCookies(c echo.Context) error {
	var cookies []*http.Cookie
	for _, cookie := range c.Cookies() {
		cookies = append(cookies, cookie)
	}
	return c.JSON(http.StatusOK, cookies)
}
