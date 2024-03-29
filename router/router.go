package router

import (
	"net/http"

	"github.com/juho05/h-bank/responses"
	"github.com/juho05/h-bank/router/middlewares"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func New() *echo.Echo {
	e := echo.New()
	e.HideBanner = true

	e.HTTPErrorHandler = responses.HandleHTTPError

	e.Pre(middleware.RemoveTrailingSlash())

	e.Use(middleware.Recover())
	e.Use(middleware.Secure())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete, http.MethodOptions},
		AllowCredentials: true,
	}))

	e.Use(middlewares.Lang)

	return e
}
