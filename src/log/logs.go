package log

import (
	log "github.com/Sirupsen/logrus"
)

func FailOnWarn(err error, msg string) {
	if err != nil {
		log.Warn(msg)
	}
}

func FailOnError(err error, msg string) {
	if err != nil {
		log.Error(msg)
	}
}

func Info(msg string) {
	log.Info(msg)
}
