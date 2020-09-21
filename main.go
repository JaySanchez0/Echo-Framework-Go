package main

import (
	"net/http"

	"github.com/labstack/echo"
)

type People struct {
	Name string
	Age  int
}

func hello(context echo.Context) error {
	return context.String(http.StatusOK, "Hola Mundo")
}

func people(context echo.Context) error {
	p := &People{Name: "juan", Age: 20}
	return context.JSON(http.StatusAccepted, p)
}

func addPeople(context echo.Context) error {
	p := new(People)
	if er := context.Bind(p); er != nil {
		return context.JSON(http.StatusBadRequest, p)
	}
	return context.JSON(http.StatusAccepted, p)
}

func getName(context echo.Context) error {
	return context.HTML(http.StatusAccepted, context.Param("name"))
}

func main() {
	p := echo.New()

	p.Static("/", "static")

	p.GET("/hello", hello)

	p.GET("/people", people)

	p.POST("/people", addPeople)

	p.GET("/people/:name", getName)

	p.Start(":80")
}
