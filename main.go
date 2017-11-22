package main

import (
	"net/http"

	"github.com/FernandoCagale/go-api-public/src/checker"
	"github.com/FernandoCagale/go-api-public/src/config"
	"github.com/FernandoCagale/go-api-public/src/handlers"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func mainPublic(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{
		"status": "ok",
	})
}

func main() {
	env := config.LoadEnv()
	app := echo.New()

	checkers := map[string]checker.Checker{
		"api":     checker.NewApi(),
		"mongodb": checker.NewMongodb(env.DatastoreURL),
	}

	app.Use(middleware.Logger())

	healthzHandler := handlers.NewHealthzHandler(checkers)
	app.GET("/health", healthzHandler.HealthzIndex)

	app.GET("/public", mainPublic)

	app.Logger.Fatal(app.Start(":" + env.Port))
}
