package logger

import (
	"errors"
	"io"
	"strings"
)

var (
	ErrInvalidFormat   = errors.New("invalid log format")
	ErrInvalidLogLevel = errors.New("invalid log level")
)

type Option interface {
	Apply(logger Log)
}

type Fields map[string]interface{}

type Level int

const (
	Debug Level = iota
	Info
	Warn
	Error
)

func (s Level) String() string {
	return levelName[s]
}

var levelName = []string{
	"debug",
	"info",
	"warn",
	"error",
}

var levelValue = map[string]Level{
	"debug": Debug,
	"info":  Info,
	"warn":  Warn,
	"error": Error,
}

func ParseLevel(l string) (Level, error) {
	v, ok := levelValue[strings.ToLower(l)]
	if !ok {
		return Debug, ErrInvalidLogLevel
	}

	return v, nil
}

type Format int

func (f Format) String() string {
	return formatName[f]
}

const (
	Json Format = iota
	Text
)

var formatName = [...]string{
	"json",
	"text",
}

var formatValue = map[string]Format{
	"json": Json,
	"text": Text,
}

func ParseFormat(f string) (Format, error) {
	v, ok := formatValue[strings.ToLower(f)]
	if !ok {
		return Json, ErrInvalidFormat
	}

	return v, nil
}

type Log interface {
	Debug(args ...interface{})
	Debugf(format string, args ...interface{})

	Info(args ...interface{})
	Infof(format string, args ...interface{})

	Warn(args ...interface{})
	Warnf(format string, args ...interface{})

	Error(args ...interface{})
	Errorf(format string, args ...interface{})

	Fatal(args ...interface{})
	Fatalf(format string, args ...interface{})

	With(fields Fields) Log

	SetLevel(level Level) error
	SetFormat(format Format) error
	SetOutput(output io.Writer)
}
