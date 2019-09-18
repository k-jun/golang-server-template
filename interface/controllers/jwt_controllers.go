package controllers

import (
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/k0kubun/pp"
	"github.com/labstack/echo"
)

const secret = "hwatXTgp1L"

// TODO Comment
func Login(c echo.Context) error {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["name"] = "FirstName LastName"
	claims["admin"] = true
	// 24 hours
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	t, err := token.SignedString([]byte(secret))
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "error Happended")
	}
	return c.JSON(http.StatusOK, map[string]string{
		"jwt_token": t,
	})
}

func MyPage(c echo.Context) error {
	user, ok := c.Get("current_user").(*jwt.Token)
	if !ok {
		return echo.NewHTTPError(http.StatusInternalServerError, "error Happended")
	}
	pp.Print(user)
	return c.JSON(http.StatusOK, user)
}

func Logout(c echo.Context) error {
	return c.String(http.StatusOK, "Deal with on client side")
}
