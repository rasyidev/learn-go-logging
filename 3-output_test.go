package learngologging

import (
	"os"
	"testing"

	"github.com/sirupsen/logrus"
)

func TestOutputFile(t *testing.T) {
	logger := logrus.New()

	file, _ := os.OpenFile("application.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	logger.SetOutput(file)

	logger.Trace("This is Trace")
	logger.Debug("This is Debug")
	logger.Info("This is Info")
	logger.Trace("This is Trace")
	logger.Error("This is Error")
}
