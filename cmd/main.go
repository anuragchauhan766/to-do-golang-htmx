package main

import (
	v1 "github.com/anuragchauhan766/to-do-golang-htmx/api/v1"
	"github.com/anuragchauhan766/to-do-golang-htmx/model"
	"github.com/anuragchauhan766/to-do-golang-htmx/utils"
	"github.com/anuragchauhan766/to-do-golang-htmx/view/pages"
	"github.com/labstack/echo/v4"
)

func main() {
	err := model.InitTodo()
	if err != nil {
		panic(err)
	}
	e := echo.New()
	e.Static("/static", "static")
	e.GET("/", func(c echo.Context) error {
		return utils.Render(c, pages.Home())
	})
	api := e.Group("/api")
	v1.ApiSetup(api)

	e.Logger.Fatal(e.Start(":5000"))
}
