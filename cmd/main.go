package main

import (
	"fmt"
	"net/http"
	"os"
	"path"
	"runtime"

	"github.com/sirupsen/logrus"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "easy")
}

func main() {
	http.HandleFunc("/", Home)

	loge := logrus.New()
	loge.SetReportCaller(true)
	loge.Formatter = &logrus.TextFormatter{
		CallerPrettyfier: func(f *runtime.Frame) (function string, file string) {
			filename := path.Base(f.File)
			return fmt.Sprintf("%s()", f.Function), fmt.Sprintf("%s,%d", filename, f.Line)
		},
		DisableColors: true,
		FullTimestamp: true,
	}
	os.MkdirAll("logs",0644)

	loge.Infof("start")
	http.ListenAndServe(":3000", nil)
}
