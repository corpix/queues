package nsq

import (
	"github.com/bitly/go-nsq"
	"github.com/sirupsen/logrus"

	"github.com/corpix/loggers"
)

const (
	LogLevelDebug   = LogLevel("debug")
	LogLevelInfo    = LogLevel("info")
	LogLevelError   = LogLevel("error")
	LogLevelWarning = LogLevel("warning")
)

var (
	NsqLogLevel = map[LogLevel]nsq.LogLevel{
		LogLevelDebug:   nsq.LogLevelDebug,
		LogLevelInfo:    nsq.LogLevelInfo,
		LogLevelError:   nsq.LogLevelError,
		LogLevelWarning: nsq.LogLevelWarning,
	}
)

type LogLevel string

func (l LogLevel) Nsq() nsq.LogLevel {
	var (
		nl nsq.LogLevel
		ok bool
	)
	nl, ok = NsqLogLevel[l]
	if !ok {
		return nsq.LogLevelInfo
	}
	return nl
}

type Logger struct {
	loggers.Logger
}

func (l *Logger) Output(_ int, s string) error {
	l.Logger.Print(s)
	return nil
}

func NewLogger(l loggers.Logger) *Logger {
	return &Logger{l}
}

func NewLogLevelFromLogrus(lv logrus.Level) LogLevel {
	switch lv {
	case logrus.DebugLevel:
		return LogLevelDebug
	case logrus.InfoLevel:
		return LogLevelInfo
	case logrus.ErrorLevel:
		return LogLevelError
	case logrus.WarnLevel:
		return LogLevelWarning
	default:
		return LogLevelInfo
	}
}
