package handler

import (
	"fmt"
	"strconv"

	"github.com/anuragchauhan766/to-do-golang-htmx/model"
	"github.com/anuragchauhan766/to-do-golang-htmx/utils"
	"github.com/anuragchauhan766/to-do-golang-htmx/view/components"
	"github.com/labstack/echo/v4"
)

func DeleteTodo(c echo.Context) error {
	todoId := c.Param("id")
	v, err := strconv.Atoi((todoId))
	if err != nil {
		return err
	}
	newTodo := model.Todos.DeleteTodo(v)

	fmt.Print(newTodo)
	return utils.Render(c, components.TodoContainer(newTodo))
}
