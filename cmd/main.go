package main

import (
	"context"

	"github.com/labstack/echo/v4"
	"github.com/ruan-molinari/ecommerce/handler"
)

type DB struct{}

func main() {
	app := echo.New()
  
	userHandler := handler.UserHandler{}
	app.Use(withUser)
	app.GET("/user", userHandler.HandleUserShow)

	app.Start(":3000")
}

func withUser(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Set("user", "a@arst.com")
		ctx := context.WithValue(c.Request().Context(), "user", "a@aa.com")
		c.SetRequest(c.Request().WithContext(ctx))
		return next(c)
	}
}
