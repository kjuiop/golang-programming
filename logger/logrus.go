package logger

import (
	"fmt"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/sirupsen/logrus"
	"io"
	"log"
	"os"
	"runtime"
	"strings"
	"time"
)

type Logrus struct {
	log *logrus.Logger
}

// LogInitialize Logger Initial
func LogInitialize(fileName, level string) (*Logrus, error) {

	l := new(Logrus)
	l.log = logrus.New()

	lv := l.getLevel(level)

	// Set Logger File Save physical
	l.log.SetFormatter(&logrus.JSONFormatter{TimestampFormat: time.RFC3339Nano})
	if len(fileName) > 0 {
		apiLogger, err := SetRollingLogFile(fileName)
		if err != nil {
			log.Printf(fileName+" : %s", err.Error())
			return nil, err
		}

		multiWriter := io.MultiWriter(os.Stdout, apiLogger)
		l.log.SetOutput(multiWriter)
	} else {
		l.log.SetOutput(os.Stdout)
	}
	l.log.SetLevel(lv)

	return l, nil
}

// SetRollingLogFile periodically changes the log file.
func SetRollingLogFile(path string) (*rotatelogs.RotateLogs, error) {
	apiLogger, err := rotatelogs.New(
		path+".%Y%m%d",
		rotatelogs.WithMaxAge(-1),
		rotatelogs.WithRotationTime(24*time.Hour),
		rotatelogs.WithLinkName(path),
	)

	if err != nil {
		return nil, err
	}

	return apiLogger, nil
}

// getLevel changes the value entered in string form to logrus.Level.
func (l *Logrus) getLevel(level string) (lv logrus.Level) {
	lv = logrus.InfoLevel
	switch strings.ToLower(level) {
	case "debug":
		lv = logrus.DebugLevel
	case "info":
		lv = logrus.InfoLevel
	case "warn":
		lv = logrus.WarnLevel
	case "error":
		lv = logrus.ErrorLevel
	default:
		logrus.Info("Unknown level string.")
	}
	return
}

func (l *Logrus) Info(prefix *logrus.Entry, format string, args ...interface{}) {
	if l.log.Level >= logrus.InfoLevel {
		prefix.Data["file"] = FileInfo(2)
		prefix.Infof(format, args...)
	}
}
func (l *Logrus) Trace(prefix *logrus.Entry, format string, args ...interface{}) {
	if l.log.Level >= logrus.TraceLevel {
		prefix.Data["file"] = FileInfo(2)
		prefix.Debugf(format, args...)
	}
}
func (l *Logrus) Debug(prefix *logrus.Entry, format string, args ...interface{}) {
	if l.log.Level >= logrus.DebugLevel {
		prefix.Data["file"] = FileInfo(2)
		prefix.Debugf(format, args...)
	}
}
func (l *Logrus) Warn(prefix *logrus.Entry, format string, args ...interface{}) {
	if l.log.Level >= logrus.WarnLevel {
		prefix.Data["file"] = FileInfo(2)
		prefix.Debugf(format, args...)
	}
}
func (l *Logrus) Error(prefix *logrus.Entry, format string, args ...interface{}) {
	if l.log.Level >= logrus.ErrorLevel {
		prefix.Data["file"] = FileInfo(2)
		prefix.Errorf(format, args...)
	}
}

func (l *Logrus) WithField(key string, obj interface{}) *logrus.Entry {
	return (*logrus.Entry)(l.log.WithField(key, obj))
}
func (l *Logrus) WithFields(fields logrus.Fields) *logrus.Entry {
	return (*logrus.Entry)(l.log.WithFields(logrus.Fields(fields)))
}

func FileInfo(skip int) string {
	_, file, line, ok := runtime.Caller(skip)
	if !ok {
		file = "<???>"
		line = 1
	} else {
		slash := strings.LastIndex(file, "/")
		if slash >= 0 {
			file = file[slash+1:]
		}
	}
	return fmt.Sprintf("%s:%d", file, line)
}
