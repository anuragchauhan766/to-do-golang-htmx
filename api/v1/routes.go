package v1

import (
	"github.com/anuragchauhan766/to-do-golang-htmx/handler"
	"github.com/labstack/echo/v4"
)

func ApiSetup(e *echo.Group) {
	api := e.Group("/v1")
	api.GET("/health", handler.HealthHandler)

	todo := api.Group("/todo")
	todo.POST("", handler.AddTodo)
	todo.DELETE("/:id", handler.DeleteTodo)
}
