package logging

import (
	"fmt"

	"github.com/sirupsen/logrus"
)

func Init() {
	l := logrus.New()
	l.SetReportCaller(true)
	
}