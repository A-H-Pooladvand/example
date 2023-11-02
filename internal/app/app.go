package app

import (
	"app/configs"
	"app/routes"
	"github.com/labstack/echo/v4"
)

func Serve() error {
	config := configs.NewApp()

	e := echo.New()
	e.HideBanner = true

	routes.Register(e)

	e.Logger.Fatal(e.Start(":" + config.Port))

	return nil
}
