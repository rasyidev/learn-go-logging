package learngologging

import (
	"testing"

	"github.com/sirupsen/logrus"
)

func TestEntry(t *testing.T) {
	logger := logrus.New()
	entry := logrus.NewEntry(logger)

	entry.WithField("username", "rasyidev").Info("This is info with username field")
}

/*
=== RUN   TestEntry
time="2022-11-07T07:31:54+07:00" level=info msg="This is info with username field" username=rasyidev
--- PASS: TestEntry (0.00s)
PASS
ok      github.com/rasyidev/learn-go-logging    0.674s
*/
