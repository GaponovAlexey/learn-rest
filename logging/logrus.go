package logging

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path"
	"runtime"

	"github.com/sirupsen/logrus"
)

type writerHook struct {
	Writer    []io.Writer
	LogLevels []logrus.Level
}

func (h *writerHook) Fire(e *logrus.Entry) error {
	line, err := e.String()
	if err != nil {
		fmt.Println(err)
	}

	for _, w := range h.Writer {
		_, err := w.Write([]byte(line))
		fmt.Println(err)
	}
	return err
}

func (hook *writerHook) levels() []logrus.Level {
	return hook.LogLevels
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
func (l *Logger) lWithFields(f map[string]interface{}) *Logger {
	return &Logger{l.WithFields(f)}
}

func Init(level string) {
	logrusLevel, err := logrus.ParseLevel(level)
	if err != nil {
		log.Fatal(err)
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

	l.AddHook(&writerHook{
		Writer:    []io.Writer{os.Stdout},
		LogLevels: logrus.AllLevels,
	})

	l.SetLevel(logrusLevel)
	e = logrus.NewEntry(l)
}
