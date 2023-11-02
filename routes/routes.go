package routes

import (
	"app/internal/handlers/users"
	"github.com/labstack/echo/v4"
)

func Register(e *echo.Echo) {
	v1 := e.Group("v1")

	userGroup := v1.Group("/users")

	userGroup.GET("", users.Index)
}
