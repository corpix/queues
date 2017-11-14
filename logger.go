package queues

import (
	"github.com/corpix/loggers"
	"github.com/corpix/loggers/logger/prefixwrapper"
)

func prefixedLogger(s string, l loggers.Logger) loggers.Logger {
	return prefixwrapper.New(
		loggerPrefix(s),
		l,
	)
}

func loggerPrefix(s string) string {
	return "[" + s + "] "
}
