package learngologging

import (
	"testing"

	"github.com/sirupsen/logrus"
)

func TestAddNewField(t *testing.T) {
	logger := logrus.New()
	logger.SetLevel(logrus.DebugLevel)

	logger.WithField("UserID", "rasyidev").Info("This is Info with additional field")
	logger.WithField("UserID", "rasyidev").
		WithField("Name", "Rasyidev").
		WithField("IsFraud", false).
		Info("This is Info with additional field")

	logger.Info("This is Info")
}

/*
=== RUN   TestAddNewField
time="2022-11-07T05:59:57+07:00" level=info msg="This is Info with additional field" UserID=rasyidev
time="2022-11-07T05:59:57+07:00" level=info msg="This is Info with additional field" IsFraud=false Name=Rasyidev UserID=rasyidev
time="2022-11-07T05:59:57+07:00" level=info msg="This is Info"
--- PASS: TestAddNewField (0.00s)
PASS
ok      github.com/rasyidev/learn-go-logging    0.780s
*/

func TestAddNewMultipleFields(t *testing.T) {
	logger := logrus.New()
	logger.SetLevel(logrus.DebugLevel)

	logger.WithField("UserID", "rasyidev").Info("This is Info with additional field")
	logger.WithFields(logrus.Fields{
		"is_fraud": false,
		"username": "Rasyidev Pro",
	})

	logger.Info("This is Info")
}

/*
=== RUN   TestAddNewMultipleFields
time="2022-11-07T06:03:51+07:00" level=info msg="This is Info with additional field" UserID=rasyidev
time="2022-11-07T06:03:51+07:00" level=info msg="This is Info"
--- PASS: TestAddNewMultipleFields (0.00s)
PASS
ok      github.com/rasyidev/learn-go-logging    0.688s
*/
