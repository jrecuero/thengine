// logger.go contains everything required for the logger used in the
// application.
package tools

import (
	"fmt"
	"os"

	"github.com/sirupsen/logrus"
)

// -----------------------------------------------------------------------------
// Private constants
// -----------------------------------------------------------------------------

const (
	logFilename = "logfile.log"
)

// -----------------------------------------------------------------------------
// Public variables
// -----------------------------------------------------------------------------

var (
	Logger *logrus.Logger
)

// -----------------------------------------------------------------------------
// Package private functions
// -----------------------------------------------------------------------------

// createLogFile function creates a new logrus.Logger instance from the given
// filename.
func createLogFile(filename string) *logrus.Logger {
	if logFile, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644); err == nil {
		logger := logrus.New()
		logger.SetOutput(logFile)
		logger.SetLevel(logrus.TraceLevel)
		return logger
	} else {
		panic(fmt.Sprintf("Error creating logfile %s Err=%+v", logFilename, err))
	}
}

// -----------------------------------------------------------------------------
// Init package
// -----------------------------------------------------------------------------

// init function provides all initialization required for the logger module.
// It initializes logging file configuration.
func init() {
	if _, err := os.Stat(logFilename); err == nil {
		fmt.Println("removing ", logFilename)
		os.Remove(logFilename)
	}
	Logger = createLogFile(logFilename)
}
