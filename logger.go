package req

import (
	"io"
	"log"
)

// Logger is the abstract logging interface, gives control to
// the Req users, choice of the logger.
type Logger interface {
	Errorf(format string, v ...any)
	Warnf(format string, v ...any)
	Debugf(format string, v ...any)
}

// NewLogger create a Logger wraps the *log.Logger
func NewLogger(output io.Writer, prefix string, flag int) Logger {
	_ = "STUB: not implemented"
	return *new(Logger)
}

func NewLoggerFromStandardLogger(l *log.Logger) Logger {
	_ = "STUB: not implemented"
	return *new(Logger)
}

func createDefaultLogger() Logger { _ = "STUB: not implemented"; return *new(Logger) }

var _ Logger = (*logger)(nil)

type disableLogger struct{}

func (l *disableLogger) Errorf(format string, v ...any) { _ = "STUB: not implemented"; return }
func (l *disableLogger) Warnf(format string, v ...any)  { _ = "STUB: not implemented"; return }
func (l *disableLogger) Debugf(format string, v ...any) { _ = "STUB: not implemented"; return }

type logger struct {
	l *log.Logger
}

func (l *logger) Errorf(format string, v ...any) { _ = "STUB: not implemented"; return }

func (l *logger) Warnf(format string, v ...any) { _ = "STUB: not implemented"; return }

func (l *logger) Debugf(format string, v ...any) { _ = "STUB: not implemented"; return }

func (l *logger) output(level, format string, v ...any) { _ = "STUB: not implemented"; return }
