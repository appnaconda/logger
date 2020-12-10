# Common Logging Interface

The aim of this log interface is to avoid doing big refactoring when changing the log library in your code. You just need to create the implementation and use the interface as your log type everywhere.

Here's how would you use logrus as your implementation:

```go
package logrus

import (
	"fmt"
	"io"

	"github.com/appnaconda/logger"
	"github.com/sirupsen/logrus"
)

type logrusLogger struct {
	logger *logrus.Entry
}

func (l *logrusLogger) Info(args ...interface{}) {
	l.logger.Info(args...)
}

func (l *logrusLogger) Infof(format string, args ...interface{}) {
	l.logger.Infof(format, args...)
}

func (l *logrusLogger) Debug(args ...interface{}) {
	l.logger.Debug(args...)
}

func (l *logrusLogger) Debugf(format string, args ...interface{}) {
	l.logger.Debugf(format, args...)
}

func (l *logrusLogger) Warn(args ...interface{}) {
	l.logger.Warn(args...)
}

func (l *logrusLogger) Warnf(format string, args ...interface{}) {
	l.logger.Warnf(format, args...)
}

func (l *logrusLogger) Error(args ...interface{}) {
	l.logger.Error(args...)
}

func (l *logrusLogger) Errorf(format string, args ...interface{}) {
	l.logger.Errorf(format, args...)
}

func (l *logrusLogger) Fatal(args ...interface{}) {
	l.logger.Fatal(args...)
}

func (l *logrusLogger) Fatalf(format string, args ...interface{}) {
	l.logger.Fatalf(format, args...)
}

func (l *logrusLogger) SetLevel(level logger.Level) error {
	var err error
	l.logger.Logger.Level, err = logrus.ParseLevel(level.String())
	return err
}

func (l *logrusLogger) SetFormat(f logger.Format) error {
	switch f {
	case logger.Text:
		l.logger.Logger.Formatter = &logrus.TextFormatter{}
	case logger.Json:
		l.logger.Logger.Formatter = &logrus.JSONFormatter{}
	default:
		return fmt.Errorf("invalid log output format: %s. JSON will be used by default", f.String())
	}

	return nil
}

func (l *logrusLogger) SetOutput(output io.Writer) {
	l.logger.Logger.Out = output
}

func (l *logrusLogger) With(fields logger.Fields) logger.Logger {
	f := logrus.Fields{}

	for k, v := range fields {
		f[k] = v
	}
	return &logrusLogger{
		l.logger.WithFields(f),
	}
}

func New(opts ...logger.Option) (logger.Logger, error) {
	logger := &logrusLogger{
		logger: logrus.NewEntry(logrus.New()),
	}

	for _, opt := range opts {
		opt.Apply(logger)
	}

	return logger, nil
}
``` 

You main program will look like this:

```go
package main

import (
    "github.com/appnaconda/logger"
    "github.com/appnaconda/logger/option"
    "github.com/your-program/logrus"
)

func main() {
  var log logger.Log

  log = logrus.New(
  	option.WithLevel(logger.DEBUG),
  	option.WithFormat(logger.JSON_FORMAT),
  )
    
  log.Debug("woohoo!!!!!")
}





```