package main

import (
	"github.com/gloomweaver/go-learning/handler"
	"github.com/gloomweaver/go-learning/view"
	"github.com/labstack/echo/v4"
)

func main() {
	app := echo.New()

	app.Static("/static", "static")

	app.GET("/", func(c echo.Context) error {
		return handler.Render(c, view.Show())
	})

	app.Logger.Fatal(app.Start(":8080"))
}
