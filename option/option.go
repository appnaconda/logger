package option

import (
	"github.com/appnaconda/logger"
)

func WithLevel(level logger.Level) logger.Option {
	return withLogLevel{level: level}
}

type withLogLevel struct {
	level logger.Level
}

func (lv withLogLevel) Apply(log logger.Log) {
	log.SetLevel(lv.level)
}

func WithFormat(format logger.Format) logger.Option {
	return withLogFormat{format: format}
}

type withLogFormat struct {
	format logger.Format
}

func (f withLogFormat) Apply(log logger.Log) {
	log.SetFormat(f.format)
}
