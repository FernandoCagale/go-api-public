package lib

import (
	"github.com/FernandoCagale/go-api-public/src/log"
	"github.com/labstack/echo"
	mgo "gopkg.in/mgo.v2"
)

func BindDb(session *mgo.Session) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			if err := verify(session); err == nil {
				c.Set("mongo", session)
			}
			return next(c)
		}
	}
}

func verify(session *mgo.Session) error {
	if err := session.Ping(); err != nil {
		session.Refresh()
		log.Info("Reconnecting mongodb!")
		if err = session.Ping(); err != nil {
			log.FailOnWarn(err, "Failed conection mongodb!")
			return err
		}
	}
	return nil
}
