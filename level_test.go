package learngologging

import (
	"testing"

	"github.com/sirupsen/logrus"
)

func TestLevel(t *testing.T) {
	logger := logrus.New()

	logger.Trace("This is Trace")
	logger.Debug("This is Debug")
	logger.Info("This is Info")
	logger.Trace("This is Trace")
	logger.Error("This is Error")

	// Kalau 2 di bawah ini dijalankan, maka program akan berhenti
	logger.Fatal("This is Fatal")
	logger.Panic("This is Panic")
}

/*
=== RUN   TestLevel
time="2022-11-06T21:43:35+07:00" level=info msg="This is Info"
time="2022-11-06T21:43:35+07:00" level=error msg="This is Error"
time="2022-11-06T21:43:35+07:00" level=fatal msg="This is Fatal"
FAIL    github.com/rasyidev/learn-go-logging    0.770s
*/

func TestSetLevel(t *testing.T) {
	logger := logrus.New()
	logger.SetLevel(logrus.TraceLevel)

	logger.Trace("This is Trace")
	logger.Debug("This is Debug")
	logger.Info("This is Info")
	logger.Trace("This is Trace")
	logger.Error("This is Error")

	// Kalau 2 di bawah ini dijalankan, maka program akan berhenti
	logger.Fatal("This is Fatal")
	logger.Panic("This is Panic")
}

/*
=== RUN   TestSetLevel
time="2022-11-06T21:44:59+07:00" level=trace msg="This is Trace"
time="2022-11-06T21:44:59+07:00" level=debug msg="This is Debug"
time="2022-11-06T21:44:59+07:00" level=info msg="This is Info"
time="2022-11-06T21:44:59+07:00" level=trace msg="This is Trace"
time="2022-11-06T21:44:59+07:00" level=error msg="This is Error"
time="2022-11-06T21:44:59+07:00" level=fatal msg="This is Fatal"
FAIL    github.com/rasyidev/learn-go-logging    0.416s
*/
