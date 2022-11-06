package learngologging

import (
	"testing"

	"github.com/sirupsen/logrus"
)

func TestOutputinJSONFormat(t *testing.T) {
	logger := logrus.New()
	logger.SetLevel(logrus.TraceLevel)
	logger.SetFormatter(&logrus.JSONFormatter{})

	logger.Trace("This is Trace")
	logger.Debug("This is Debug")
	logger.Info("This is Info")
	logger.Trace("This is Trace")
	logger.Error("This is Error")
}

/*
=== RUN   TestOutputinJSONFormat
{"level":"trace","msg":"This is Trace","time":"2022-11-06T23:40:51+07:00"}
{"level":"debug","msg":"This is Debug","time":"2022-11-06T23:40:51+07:00"}
{"level":"info","msg":"This is Info","time":"2022-11-06T23:40:51+07:00"}
{"level":"trace","msg":"This is Trace","time":"2022-11-06T23:40:51+07:00"}
{"level":"error","msg":"This is Error","time":"2022-11-06T23:40:51+07:00"}
--- PASS: TestOutputinJSONFormat (0.00s)
PASS
ok      github.com/rasyidev/learn-go-logging    0.388s
*/
