package logging

import (
	"fmt"
	"gin-boilerplate/configs"
	"io"
	"os"
	"runtime"
	"strings"

	"github.com/sirupsen/logrus"
)

var Logger = logrus.New()

                                                                                                                                                                               


func init() {
	Logger.SetReportCaller(true)
	Logger.Out = os.Stdout

	if configs.Config.LogFormat == "text" {

		formatter := &logrus.TextFormatter{

			TimestampFormat:        "01-01-2006 15:04:05", // Golang's time formatting insanity is to use so called "reference date" as a template so instead of "YYYY-MM-DD" one have to use "2006-01-02" literally. 
			FullTimestamp:          true,
			DisableLevelTruncation: true,
			CallerPrettyfier: func(f *runtime.Frame) (string, string) {

				return "", fmt.Sprintf(" file-> %s line-> %d function -> %s", GetMainFile(f.File), f.Line, f.Function)
			},
		}
		Logger.SetFormatter(formatter)

	} else {
		formatter := &logrus.JSONFormatter{

			TimestampFormat: "01-01-2006 15:04:05",
			CallerPrettyfier: func(f *runtime.Frame) (string, string) {

				return FuncName(f.Function), fmt.Sprintf("%s, line %d", GetMainFile(f.File), f.Line)
			},
		}
		Logger.SetFormatter(formatter)
	}
	if configs.Config.WriteLog {
		WriteLogOut()
	}
}

func GetMainFile(fullpath string) string {
	result := strings.Split(fullpath, "/")

	return result[len(result)-1] // getting the last file where logging being writtem
}

func FuncName(function string) string {

	s := strings.Split(function, ".")
	funcname := s[len(s)-1]
	return funcname
}

func WriteLogOut() {

	f, err := os.OpenFile(configs.Config.LogFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		Logger.Fatalf("error opening file: %v", err)
	}
	wrt := io.MultiWriter(os.Stdout, f)

	Logger.SetOutput(wrt)
}
