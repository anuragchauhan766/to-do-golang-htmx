package handler

import (
	"fmt"

	"github.com/anuragchauhan766/to-do-golang-htmx/model"
	"github.com/anuragchauhan766/to-do-golang-htmx/utils"
	"github.com/anuragchauhan766/to-do-golang-htmx/view/components"
	"github.com/labstack/echo/v4"
)

func AddTodo(c echo.Context) error {
	todoContent := c.FormValue("todoContent")
	newTodo := model.Todos.AddTodo(todoContent)
	fmt.Print(newTodo)
	return utils.Render(c, components.TodoContainer(newTodo))
}
