package logging

import (
	log "github.com/sirupsen/logrus"
	"os"
)

//I did everything through Entry so that later
//if I want to change WithFields , I don't have to create a new logger
var e *log.Entry

type Logger struct {
	*log.Entry
}

func GetLogger() Logger {
	return Logger{e}
}
func init() {
	l := log.New()
	l.SetReportCaller(true)
	l.SetOutput(os.Stdout)
	l.SetFormatter(&log.JSONFormatter{
		TimestampFormat:   "2006-01-02 15:04:05",
		DisableTimestamp:  false,
		DisableHTMLEscape: false,
		DataKey:           "",
		CallerPrettyfier:  nil,
		PrettyPrint:       true,
	})
	e = log.NewEntry(l)
}
