package logger

import (
	"github.com/astota/go-logging"
	"github.com/sirupsen/logrus"
	"os"
	"reflect"
)

// logger is global logger instance
var logger *logrus.Logger

// defaultLogger logger instance
type defaultLogger struct {
	log *logrus.Entry
}

// AddFields adds fields to logger backend
func (l *defaultLogger) AddFields(fields logging.Fields) logging.Logger {
	f := logrus.Fields{}
	for key, val := range fields {
		f[key] = val
	}

	return &defaultLogger{log: l.log.WithFields(f)}
}

// AddFields adds fields to logger backend
func (l *defaultLogger) AddFieldsToCurrent(fields logging.Fields) logging.Logger {
	f := logrus.Fields{}
	for key, val := range fields {
		f[key] = val
	}
	l.log = l.log.WithFields(f)

	return &defaultLogger{log: l.log}
}

//SetLevel adds logging level
func (l *defaultLogger) SetLevel(lvl logging.Level) {
	switch lvl {
	case logging.LevelDebug:
		logger.SetLevel(logrus.DebugLevel)
	case logging.LevelInfo:
		logger.SetLevel(logrus.InfoLevel)
	case logging.LevelError:
		logger.SetLevel(logrus.ErrorLevel)
	case logging.LevelFatal:
		logger.SetLevel(logrus.FatalLevel)
	default:
		// Not tested because this should never happen
		// It is just safeguard
		logger.SetLevel(logrus.InfoLevel)
	}
}

// Debug writes debug level log entry
func (l *defaultLogger) Debug(str string) {
	l.log.Debug(str)
}

// Debugf writes debug level log entry with formatting support
func (l *defaultLogger) Debugf(f string, args ...interface{}) {
	l.log.Debugf(f, args...)
}

// Info writes info level log entry
func (l *defaultLogger) Info(str string) {
	l.log.Info(str)
}

// Infof writes info level log entry with formatting support
func (l *defaultLogger) Infof(f string, args ...interface{}) {
	l.log.Infof(f, args...)
}

// Error writes error level log entry
func (l *defaultLogger) Error(str string) {
	l.log.Error(str)
}

// Errorf writes error level log entry with formatting support
func (l *defaultLogger) Errorf(f string, args ...interface{}) {
	l.log.Errorf(f, args...)
}

// Fatal writes fatal level log entry
func (l *defaultLogger) Fatal(str string) {
	// Not tested, this will call Exit(1), which will end
	// test execution.
	l.log.Fatal(str)
}

// Fatalf writes fatal level log entry with formatting support
func (l *defaultLogger) Fatalf(f string, args ...interface{}) {
	// Not tested, this will call Exit(1), which will end
	// test execution.
	l.log.Fatalf(f, args...)
}

func (l *defaultLogger) WithError(err error) logging.Logger {
	if err == nil {
		return l
	}
	f := logging.Fields{}
	f["error.message"] = err.Error()

	if t := reflect.TypeOf(err); t.Kind() == reflect.Ptr {
		f["error.kind"] = t.Elem().Name()
	} else {
		f["error.kind"] = t.Name()
	}

	return l.AddFields(f)
}

// newLogger will create new default logger instance
func newLogger() logging.Logger {
	return &defaultLogger{log: logrus.NewEntry(logger)}
}

func init() {
	logging.Register("default-logger", newLogger)
	logger = logrus.New()
	logger.Formatter = &logrus.JSONFormatter{
		FieldMap: logrus.FieldMap{
			logrus.FieldKeyTime: "timestamp",
			logrus.FieldKeyMsg:  "message",
		},
	}
	logger.Out = os.Stdout
}
