package common

import "github.com/sirupsen/logrus"

func newLogger(module string) *logrus.Entry {
	logrus.SetFormatter(&logrus.JSONFormatter{})
	return logrus.WithFields(logrus.Fields{"app": module})
}

func WithLogger(module string) *logrus.Entry {
	return Logger.WithFields(logrus.Fields{"module": module})
}

var Logger = newLogger("scrapy")
