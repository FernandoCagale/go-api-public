package main

import (
	"time"

	"github.com/FernandoCagale/go-api-public/src/checker"
	"github.com/FernandoCagale/go-api-public/src/config"
	"github.com/FernandoCagale/go-api-public/src/datastore"
	"github.com/FernandoCagale/go-api-public/src/handlers"
	"github.com/FernandoCagale/go-api-public/src/lib"
	"github.com/FernandoCagale/go-api-public/src/log"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	mgo "gopkg.in/mgo.v2"
)

func main() {
	var session *mgo.Session
	env := config.LoadEnv()
	app := echo.New()

	go bindDatastore(app, session, env.DatastoreURL)

	checkers := map[string]checker.Checker{
		"api":     checker.NewApi(),
		"mongodb": checker.NewMongodb(env.DatastoreURL),
	}

	defer session.Close()

	app.Use(middleware.Logger())

	healthzHandler := handlers.NewHealthzHandler(checkers)
	app.GET("/health", healthzHandler.HealthzIndex)

	group := app.Group("/v1")

	projectHandler := handlers.NewProjectHandler()

	group.GET("/project", projectHandler.GetAllProject)
	group.POST("/project", projectHandler.SaveProject)
	group.GET("/project/:id", projectHandler.GetProject)
	group.PUT("/project/:id", projectHandler.UpdateProject)
	group.DELETE("/project/:id", projectHandler.DeleteProject)

	app.Logger.Fatal(app.Start(":" + env.Port))
}

func bindDatastore(app *echo.Echo, session *mgo.Session, url string) {
	for {
		session, err := datastore.New(url)
		log.FailOnWarn(err, "Failed to init mongodb connection!")
		if err == nil {
			app.Use(lib.BindDb(session))
			break
		}
		time.Sleep(time.Second * 5)
	}
}
