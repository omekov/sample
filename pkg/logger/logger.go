package logger

import "github.com/sirupsen/logrus"

func Get() *logrus.Logger {
	return logrus.New()
}
