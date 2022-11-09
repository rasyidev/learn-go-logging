package learngologging

import (
	"fmt"
	"testing"

	"github.com/sirupsen/logrus"
)

type SlackLog struct{}

func (sl *SlackLog) Levels() []logrus.Level {
	return []logrus.Level{logrus.ErrorLevel, logrus.WarnLevel}
}

func (sl *SlackLog) Fire(entry *logrus.Entry) error {
	// eksekusi log
	fmt.Println("Slack Log dijalankan: ", entry.Level, entry.Message)
	return nil
}

func TestHook(t *testing.T) {
	logger := logrus.New()
	logger.AddHook(&SlackLog{})

	logger.Info("Ini adalah log info")
	logger.Error("Ini adalah log Error")
	logger.Warn("Ini adalah log Warn")
}

/*
=== RUN   TestHook
time="2022-11-09T12:14:23+07:00" level=info msg="Ini adalah log info"
Slack Log dijalankan:  error Ini adalah log Error
time="2022-11-09T12:14:23+07:00" level=error msg="Ini adalah log Error"
Slack Log dijalankan:  warning Ini adalah log Warn
time="2022-11-09T12:14:23+07:00" level=warning msg="Ini adalah log Warn"
--- PASS: TestHook (0.00s)
PASS
ok      github.com/rasyidev/learn-go-logging    0.356s
*/
