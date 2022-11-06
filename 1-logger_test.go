package learngologging

import (
	"fmt"
	"log"
	"testing"

	"github.com/sirupsen/logrus"
)

func TestLogger(t *testing.T) {
	logger := logrus.New()

	fmt.Println("Hello logger using fmt")
	log.Println("Hello Logger using log")
	logger.Println("Hello Logger using logrus")
}

/*
=== RUN   TestLogger
Hello logger using fmt
2022/11/06 21:34:47 Hello Logger using log
time="2022-11-06T21:34:47+07:00" level=info msg="Hello Logger using logrus"
--- PASS: TestLogger (0.00s)
PASS
ok      github.com/rasyidev/learn-go-logging    0.340s
*/
