package nsq

import (
	"github.com/bitly/go-nsq"
	"github.com/sirupsen/logrus"

	"github.com/corpix/logger"
)

//

const (
	LogLevelDebug   = nsq.LogLevelDebug
	LogLevelInfo    = nsq.LogLevelInfo
	LogLevelError   = nsq.LogLevelError
	LogLevelWarning = nsq.LogLevelWarning
)

//

type Logger struct {
	logger.Logger
}

func (l *Logger) Output(_ int, s string) error {
	l.Logger.Print(s)
	return nil
}

//

func NewLogger(l logger.Logger) *Logger {
	return &Logger{l}
}

//

func NewLogLevelFromLogrus(lv logrus.Level) nsq.LogLevel {
	switch lv {
	case logrus.DebugLevel:
		return nsq.LogLevelDebug
	case logrus.InfoLevel:
		return nsq.LogLevelInfo
	case logrus.ErrorLevel:
		return nsq.LogLevelError
	case logrus.WarnLevel:
		return nsq.LogLevelWarning
	default:
		return nsq.LogLevelInfo
	}
}
