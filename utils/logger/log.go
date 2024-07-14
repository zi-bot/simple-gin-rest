package logger

import (
	"sync"

	"github.com/sirupsen/logrus"
)

var (
	once     sync.Once
	instance *logrus.Logger
)

// GetLogger returns the singleton instance of the logger
func GetLogger() *logrus.Logger {
	once.Do(func() {
		instance = logrus.New()
		// Customize logger here if needed, for example:
		instance.SetFormatter(&logrus.JSONFormatter{})
		instance.SetLevel(logrus.InfoLevel)
	})
	return instance
}
