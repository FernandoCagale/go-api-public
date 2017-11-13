package main

import (
	"net/http"

	"github.com/labstack/echo"
)

func mainPublic(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{
		"status": "ok",
	})
}

func main() {
	e := echo.New()

	e.GET("/public", mainPublic)

	e.Logger.Fatal(e.Start(":8000"))
}
