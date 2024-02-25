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
// Init package
// -----------------------------------------------------------------------------

// init function provides all initialization required for the logger module.
// It initializes logging file configuration.
func init() {
	if _, err := os.Stat(logFilename); err != nil {
		os.Remove(logFilename)
		if logFile, err := os.OpenFile(logFilename, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644); err == nil {
			Logger = logrus.New()
			Logger.SetOutput(logFile)
			Logger.SetLevel(logrus.TraceLevel)
		} else {
			panic(fmt.Sprintf("Error creating logfile %s Err=%+v", logFilename, err))
		}
	}
}
