package logging

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path"
	"runtime"

	"github.com/sirupsen/logrus"
)

type writeHook struct {
	Writer    []io.Writer
	LogLevels []logrus.Level
}

func (h *writeHook) Fire(e *logrus.Entry) error {
	line, err := e.String()
	if err != nil {
		fmt.Println(err)
	}
	
	for _, w := range h.Writer {
		_, err := w.Write([]byte(line))
	}
	return err
}

func (h *writeHook) levels() []logrus.Level {
	return h.LogLevels
}

var e *logrus.Entry

type Logger struct {
	*logrus.Entry
}

func GetLogger() *Logger {
	return &Logger{e}
}

func (l *Logger) lWithField(k string, v interface{}) *Logger {
	return &Logger{l.WithField(k, v)}
}
func (l *Logger) lWithField(f map[string]interface{}) *Logger {
	return &Logger{l.WithFields(f)}
}

func Init() {
	logrusLevel, err := logrus.ParseLevel(level)
	if err != nil {
		fmt.Println(err)
	}

	l := logrus.New()
	l.SetReportCaller(true)
	l.Formatter = &logrus.TextFormatter{
		CallerPrettyfier: func(f *runtime.Frame) (string, string) {
			filename := path.Base(f.File)
			return fmt.Sprintf("%v %v", filename, f.Line), fmt.Sprintf("%s", f.Function)
		},
		DisableColors: false,
		FullTimestamp: true,
	}
	l.SetOutput(ioutil.Discard)
	l.AddHook(&writeHook{
		Writer:    []io.Writer{os.Stdout},
		LogLevels: logrus.AllLevels,
	})

	l.SetLevel(logrus.TraceLevel)
	e = logrus.NewEntry(l)
}
