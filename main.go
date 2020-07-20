package main

import (
	"github.com/labstack/echo/v4"
	"github.com/pgmot/superwebapp/pkg/hello"
	"net/http"
)

func main() {
	e := echo.New()
	e.GET("/hello/:name", func(c echo.Context) error {
		name := c.Param("name")
		return c.String(http.StatusOK, hello.Hello(name))
	})
	e.Logger.Fatal(e.Start(":1323"))
}
