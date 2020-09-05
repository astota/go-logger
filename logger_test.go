package logger

import (
	"github.com/astota/go-logging"
	"bytes"
	"github.com/sirupsen/logrus"
	"reflect"
	"regexp"
	"strings"
	"testing"
)

func TestLogrusImplementsLogging(t *testing.T) {
	var log logging.Logger

	log = newLogger()
	if reflect.ValueOf(log).Type().String() != "*logger.defaultLogger" {
		t.Errorf("Invalid return type, expected: %s, got: %s",
			"*logger.defaultLogger", reflect.ValueOf(log).Type().String())
	}
}

func TestDebug(t *testing.T) {
	tests := []struct {
		name     string
		params   string
		expected *regexp.Regexp
	}{
		{"Simple line", "test", regexp.MustCompile(`^{"level":"debug","message":"test","timestamp":"[0-9+-: TZ]*"}$`)},
	}

	logging.UseLogger("default-logger")

	for _, tst := range tests {
		buffer := &bytes.Buffer{}
		logger.Out = &TestWriter{buffer: buffer}
		log := logging.NewLogger()
		log.SetLevel(logging.LevelDebug)
		log.Debug(tst.params)
		bf := strings.TrimSpace(buffer.String())
		if !tst.expected.MatchString(bf) {
			t.Errorf("Invalid output, expected: '%s' got: '%s'",
				tst.expected.String(), bf)
		}
	}
}

func TestDebugf(t *testing.T) {
	tests := []struct {
		name     string
		format   string
		params   interface{}
		expected *regexp.Regexp
	}{
		{"Simple line", "val=%2.2f", 12.345, regexp.MustCompile(`^{"level":"debug","message":"val=12.35","timestamp":"[0-9+-: TZ]*"}$`)},
	}

	logging.UseLogger("default-logger")

	for _, tst := range tests {
		buffer := &bytes.Buffer{}
		logger.Out = &TestWriter{buffer: buffer}
		log := logging.NewLogger()
		log.SetLevel(logging.LevelDebug)
		log.Debugf(tst.format, tst.params)
		bf := strings.TrimSpace(buffer.String())
		if !tst.expected.MatchString(bf) {
			t.Errorf("Invalid output, expected: '%s' got: '%s'",
				tst.expected.String(), bf)
		}
	}
}

func TestInfo(t *testing.T) {
	tests := []struct {
		name     string
		params   string
		expected *regexp.Regexp
	}{
		{"Simple line", "test", regexp.MustCompile(`^{"level":"info","message":"test","timestamp":"[0-9+-: TZ]*"}$`)},
	}

	logging.UseLogger("default-logger")

	for _, tst := range tests {
		buffer := &bytes.Buffer{}
		logger.Out = &TestWriter{buffer: buffer}
		log := logging.NewLogger()
		log.SetLevel(logging.LevelDebug)
		log.Info(tst.params)
		bf := strings.TrimSpace(buffer.String())
		if !tst.expected.MatchString(bf) {
			t.Errorf("Invalid output, expected: '%s' got: '%s'",
				tst.expected.String(), bf)
		}
	}
}

func TestInfof(t *testing.T) {
	tests := []struct {
		name     string
		format   string
		params   interface{}
		expected *regexp.Regexp
	}{
		{"Simple line", "val=%2.2f", 12.345, regexp.MustCompile(`^{"level":"info","message":"val=12.35","timestamp":"[0-9+-: TZ]*"}$`)},
	}

	logging.UseLogger("default-logger")

	for _, tst := range tests {
		buffer := &bytes.Buffer{}
		logger.Out = &TestWriter{buffer: buffer}
		log := logging.NewLogger()
		log.SetLevel(logging.LevelDebug)
		log.Infof(tst.format, tst.params)
		bf := strings.TrimSpace(buffer.String())
		if !tst.expected.MatchString(bf) {
			t.Errorf("Invalid output, expected: '%s' got: '%s'",
				tst.expected.String(), bf)
		}
	}
}

func TestError(t *testing.T) {
	tests := []struct {
		name     string
		params   string
		expected *regexp.Regexp
	}{
		{"Simple line", "test", regexp.MustCompile(`^{"level":"error","message":"test","timestamp":"[0-9+-: TZ]*"}$`)},
	}

	logging.UseLogger("default-logger")

	for _, tst := range tests {
		buffer := &bytes.Buffer{}
		logger.Out = &TestWriter{buffer: buffer}
		log := logging.NewLogger()
		log.SetLevel(logging.LevelDebug)
		log.Error(tst.params)
		bf := strings.TrimSpace(buffer.String())
		if !tst.expected.MatchString(bf) {
			t.Errorf("Invalid output, expected: '%s' got: '%s'",
				tst.expected.String(), bf)
		}
	}
}

func TestErrorf(t *testing.T) {
	tests := []struct {
		name     string
		format   string
		params   interface{}
		expected *regexp.Regexp
	}{
		{"Simple line", "val=%2.2f", 12.345, regexp.MustCompile(`^{"level":"error","message":"val=12.35","timestamp":"[0-9+-: TZ]*"}$`)},
	}

	logging.UseLogger("default-logger")

	for _, tst := range tests {
		buffer := &bytes.Buffer{}
		logger.Out = &TestWriter{buffer: buffer}
		log := logging.NewLogger()
		log.SetLevel(logging.LevelDebug)
		log.Errorf(tst.format, tst.params)
		bf := strings.TrimSpace(buffer.String())
		if !tst.expected.MatchString(bf) {
			t.Errorf("Invalid output, expected: '%s' got: '%s'",
				tst.expected.String(), bf)
		}
	}
}

func TestLoggingLevels(t *testing.T) {
	tests := []struct {
		name     string
		lvl      logging.Level
		expected logrus.Level
	}{
		{"debug", logging.LevelDebug, logrus.DebugLevel},
		{"info", logging.LevelInfo, logrus.InfoLevel},
		{"error", logging.LevelError, logrus.ErrorLevel},
		{"fatal", logging.LevelFatal, logrus.FatalLevel},
	}

	logging.UseLogger("default-logger")
	log := logging.NewLogger()
	for _, tst := range tests {
		log.SetLevel(tst.lvl)
		fl := log.(*defaultLogger)
		if fl.log.Logger.Level != tst.expected {
			t.Errorf("Invalid logging level (%s), expected: %v, got: %v", tst.name, tst.expected, fl.log.Logger.Level)
		}
	}
}

func TestAddFields(t *testing.T) {
	tests := []struct {
		name     string
		params   logging.Fields
		expected *regexp.Regexp
	}{
		{"Simple line", logging.Fields{"test": "test!"}, regexp.MustCompile(`^{"level":"info","message":"fields","test":"test!","timestamp":"[0-9+-: TZ]*"}$`)},
	}

	logging.UseLogger("default-logger")

	for _, tst := range tests {
		buffer := &bytes.Buffer{}
		logger.Out = &TestWriter{buffer: buffer}
		log := logging.NewLogger()
		log.SetLevel(logging.LevelDebug)
		log = log.AddFields(tst.params)
		log.Info("fields")
		bf := strings.TrimSpace(buffer.String())
		if !tst.expected.MatchString(bf) {
			t.Errorf("Invalid output, expected: '%s' got: '%s'",
				tst.expected.String(), bf)
		}
	}
}

func TestTwoLoggerAndFieldsToCurrent(t *testing.T) {
	logging.UseLogger("default-logger")

	buffer := &bytes.Buffer{}
	logger.Out = &TestWriter{buffer: buffer}
	logger1 := logging.NewLogger()
	logger2 := logging.NewLogger()
	logger3 := logger1.AddFieldsToCurrent(logging.Fields{
		"logger1": "ok",
	})
	logger4 := logging.NewLogger()

	logger1.Info("logger1")
	logger2.Info("logger2")
	logger3.Info("logger3")
	logger4.Info("logger4")

	ex1 := regexp.MustCompile(`{"level":"info","logger1":"ok","message":"logger1","timestamp":"[0-9+-: TZ]*"}`)
	ex2 := regexp.MustCompile(`{"level":"info","message":"logger2","timestamp":"[0-9+-: TZ]*"}`)
	ex3 := regexp.MustCompile(`{"level":"info","logger1":"ok","message":"logger3","timestamp":"[0-9+-: TZ]*"}`)
	ex4 := regexp.MustCompile(`{"level":"info","message":"logger4","timestamp":"[0-9+-: TZ]*"}`)

	bf := strings.TrimSpace(buffer.String())
	if !ex1.MatchString(bf) {
		t.Errorf("Invalid output, expected: '%s' got: '%s'",
			ex1.String(), bf)
	}
	if !ex2.MatchString(bf) {
		t.Errorf("Invalid output, expected: '%s' got: '%s'",
			ex2.String(), bf)
	}
	if !ex3.MatchString(bf) {
		t.Errorf("Invalid output, expected: '%s' got: '%s'",
			ex3.String(), bf)
	}
	if !ex4.MatchString(bf) {
		t.Errorf("Invalid output, expected: '%s' got: '%s'",
			ex4.String(), bf)
	}
}

func TestTwoLoggerAndFields(t *testing.T) {
	logging.UseLogger("default-logger")

	buffer := &bytes.Buffer{}
	logger.Out = &TestWriter{buffer: buffer}
	logger1 := logging.NewLogger()
	logger2 := logging.NewLogger()
	logger3 := logger1.AddFields(logging.Fields{
		"logger3": "ok",
	})
	logger4 := logging.NewLogger()

	logger1.Info("logger1")
	logger2.Info("logger2")
	logger3.Info("logger3")
	logger4.Info("logger4")

	ex1 := regexp.MustCompile(`{"level":"info","message":"logger1","timestamp":"[0-9+-: TZ]*"}`)
	ex2 := regexp.MustCompile(`{"level":"info","message":"logger2","timestamp":"[0-9+-: TZ]*"}`)
	ex3 := regexp.MustCompile(`{"level":"info","logger3":"ok","message":"logger3","timestamp":"[0-9+-: TZ]*"}`)
	ex4 := regexp.MustCompile(`{"level":"info","message":"logger4","timestamp":"[0-9+-: TZ]*"}`)

	bf := strings.TrimSpace(buffer.String())
	if !ex1.MatchString(bf) {
		t.Errorf("Invalid output, expected: '%s' got: '%s'",
			ex1.String(), bf)
	}
	if !ex2.MatchString(bf) {
		t.Errorf("Invalid output, expected: '%s' got: '%s'",
			ex2.String(), bf)
	}
	if !ex3.MatchString(bf) {
		t.Errorf("Invalid output, expected: '%s' got: '%s'",
			ex3.String(), bf)
	}
	if !ex4.MatchString(bf) {
		t.Errorf("Invalid output, expected: '%s' got: '%s'",
			ex4.String(), bf)
	}
}

type testError struct {
	msg string
}

func (e testError) Error() string {
	return e.msg
}

func TestWithError(t *testing.T) {
	tests := []struct {
		name string
		err  error
		ex   *regexp.Regexp
	}{
		{"struct", testError{msg: "struct"}, regexp.MustCompile(`{"error.kind":"testError","error.message":"struct","level":"info","message":"test","timestamp":"[0-9+-: TZ]*"}`)},
		{"pointer", &testError{msg: "pointer"}, regexp.MustCompile(`{"error.kind":"testError","error.message":"pointer","level":"info","message":"test","timestamp":"[0-9+-: TZ]*"}`)},
		{"nil", nil, regexp.MustCompile(`{"level":"info","message":"test","timestamp":"[0-9+-: TZ]*"}`)},
	}

	for _, tst := range tests {
		t.Run(tst.name, func(t *testing.T) {
			logging.UseLogger("default-logger")

			buffer := &bytes.Buffer{}
			logger.Out = &TestWriter{buffer: buffer}
			l := logging.NewLogger()

			l = l.WithError(tst.err)
			l.Info("test")

			bf := strings.TrimSpace(buffer.String())
			if !tst.ex.MatchString(bf) {
				t.Errorf("Invalid output, expected: '%s' got: '%s'",
					tst.ex.String(), bf)
			}
		})
	}
}

type TestWriter struct {
	buffer *bytes.Buffer
}

func (tw *TestWriter) Write(p []byte) (n int, err error) {
	return tw.buffer.Write(p)
}
