package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/labstack/echo"
)

func HelloEcho(c echo.Context) error {
	return c.String(http.StatusOK, "this api is working! Ver 1.0.0")
}

func CSRF(c echo.Context) error {
	token, ok := c.Get("csrf").(string)
	if !ok {
		// TODO change message
		return echo.NewHTTPError(http.StatusInternalServerError, "error Happended")
	}
	return c.String(http.StatusOK, token)
}

type timeoutController struct{}

type timeoutBody struct {
	Str            string  `json:"str"`
	Int            int64   `json:"int"`
	Float          float64 `json:"float"`
	WithUnderScore string  `json:"with_under_score"`
}

func (h *timeoutController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// <-time.After(1 * time.Minute)
	b := timeoutBody{Str: "string", Int: 1, Float: 1.1, WithUnderScore: "be careful to underscore"}
	bb, err := json.Marshal(b)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(bb)

	if err != nil {
		log.Println("err:", err)
	}
}

func Timeout(c echo.Context) error {
	var tc timeoutController
	timeoutLimit := 3 * time.Second
	// TODO to customize this handler, need original timeoutHandler
	http.TimeoutHandler(&tc, timeoutLimit, "timeout error message").ServeHTTP(c.Response(), c.Request())
	return nil
}
