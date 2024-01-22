package logger

import (
	"os"

	log "github.com/sirupsen/logrus"
)

func Start() {
	log.SetFormatter(&log.JSONFormatter{})

	log.SetOutput(os.Stdout)

	log.SetLevel(log.WarnLevel)
}

func Info(args interface{}) {
	log.WithFields(log.Fields{
		"animal": "walrus",
		"size":   10,
	}).Info(args)
}

func Warn(args interface{}) {
	log.WithFields(log.Fields{
		"omg":    true,
		"number": 122,
	}).Warn(args)
}

func Fatal(args interface{}) {
	log.WithFields(log.Fields{
		"omg":    true,
		"number": 100,
	}).Fatal(args)
}
