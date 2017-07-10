package logrus

import (
	"github.com/sirupsen/logrus"

	"github.com/corpix/logger"
)

// Logrus is a logrus for logger that implements
// io.Writer interface.
type Logrus struct {
	*logrus.Logger
}

// Write slice of bytes into the logger and return number of written
// bytes and error value of present.
func (l *Logrus) Write(buf []byte) (int, error) {
	n := len(buf) + 1
	l.Printf("%s\n", buf)
	return n, nil
}

// Level returns a current logger level number.
func (l *Logrus) Level() interface{} {
	return l.Logger.Level
}

// New wraps logrus logger with binding.
func New(l *logrus.Logger) logger.Logger {
	return &Logrus{l}
}
