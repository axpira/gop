package log

import "context"

var (
	// DefaultLogger logger implementation to use
	DefaultLogger Logger
)

type LoggerOption interface {
	Update(Logger) Logger
}

type LoggerOptionFunc func(Logger) Logger

func (l LoggerOptionFunc) Update(lo Logger) Logger {
	return l(lo)
}

type Logger interface {
	// Level returns the highest level that is enabled
	Level() Level
	// HasLevel checks if the level is enabled
	HasLevel(Level) bool
	// With create a new Logger with opts informed
	With(...LoggerOption) Logger
	// NewFieldBuilder create a new FieldBuilder
	NewFieldBuilder() FieldBuilder

	// Log send a log message with level and fields if the level is enabled
	Log(Level, FieldBuilder)

	// Trc send a log message with level trace and fields
	Trc(FieldBuilder)
	// Trace send a log message with level trace and message
	Trace(string)
	// Tracef send a log message with level trace and message formatted
	Tracef(string, ...interface{})

	// Dbg send a log message with level debug and fields
	Dbg(FieldBuilder)
	// Debug send a log message with level debug and message
	Debug(string)
	// Debugf send a log message with level debug and message formatted
	Debugf(string, ...interface{})

	// Inf send a log message with level info and fields
	Inf(FieldBuilder)
	// Info send a log message with level info and message
	Info(string)
	// Infof send a log message with level info and message formatted
	Infof(string, ...interface{})

	// Wrn send a log message with level warning and fields
	Wrn(FieldBuilder)
	// Warn send a log message with level warning and message
	Warn(string)
	// Warnf send a log message with level warning and message formatted
	Warnf(string, ...interface{})

	// Err send a log message with level error and fields
	Err(FieldBuilder)
	// Error send a log message with level error, message and error
	Error(string, error)
	// Errof send a log message with level error and message formatted
	Errorf(string, ...interface{})

	// Ftl send a log message with level fatal and fields
	Ftl(FieldBuilder)
	// Fatal send a log message with level fatal and message
	Fatal(string)
	// Fatalf send a log message with level fatal and message formatted
	Fatalf(string, ...interface{})

	// Pnc send a log message with level panic and fields
	Pnc(FieldBuilder)
	// Panic send a log message with level panic and message
	Panic(string)
	// Panicf send a log message with level panic and message formatted
	Panicf(string, ...interface{})

	Print(...interface{})
	Printf(string, ...interface{})
	Println(...interface{})

	// Write implements io.Writer
	Write([]byte) (int, error)

	// FromCtx restore a logger from context
	FromCtx(context.Context) Logger

	// ToCtx store a logger in context
	ToCtx(context.Context) context.Context
}

// With create a new Logger with opts updated
func With(opts ...LoggerOption) Logger {
	return DefaultLogger.With(opts...)
}

// Log send a log message using DefaultLogger with sent level and fields
func Log(level Level, fieldBuilder FieldBuilder) {
	DefaultLogger.Log(level, fieldBuilder)
}

// Trc send a log message using DefaultLogger with level trace and fields
func Trc(fieldBuilder FieldBuilder) {
	DefaultLogger.Trc(fieldBuilder)
}

// Trace send a log message using DefaultLogger with level info and message
func Trace(msg string) {
	DefaultLogger.Trace(msg)
}

// Tracef send a log message using DefaultLogger with level info and message formatted
func Tracef(format string, args ...interface{}) {
	DefaultLogger.Tracef(format, args...)
}

// Inf send a log message using DefaultLogger with level info and fields
func Inf(fieldBuilder FieldBuilder) {
	DefaultLogger.Inf(fieldBuilder)
}

// Info send a log message using DefaultLogger with level info and message
func Info(msg string) {
	DefaultLogger.Info(msg)
}

// Infof send a log message using DefaultLogger with level info and message formatted
func Infof(format string, args ...interface{}) {
	DefaultLogger.Infof(format, args...)
}

// Err send a log message using DefaultLogger with level error and fields
func Err(fieldBuilder FieldBuilder) {
	DefaultLogger.Err(fieldBuilder)
}

// Error send a log message using DefaultLogger with level error and message
func Error(msg string, err error) {
	DefaultLogger.Error(msg, err)
}

// Errorf send a log message using DefaultLogger with level error and message formatted
func Errorf(format string, args ...interface{}) {
	DefaultLogger.Errorf(format, args...)
}

// FromCtx restore Logger from a context
func FromCtx(ctx context.Context) Logger {
	return DefaultLogger.FromCtx(ctx)
}
