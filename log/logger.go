package log

import (
	"io"
	"time"
)

const (
	panicMessage = "logger must be set"
)

type Level uint8

const (
	NoLevel Level = iota
	LevelTrace
	LevelDebug
	LevelInfo
	LevelWarn
	LevelError
	LevelFatal
	LevelPanic
	Disabled
)

var (
	defaultLogger Logger = emptyLogger{}
)

type LevelLogger interface {
	Trace(string)
	Debug(string)
	Info(string)
	Warn(string)
	Error(string)
	Fatal(string)
	Panic(string)

	Write([]byte) (int, error)
}

type Event interface {
	LevelLogger
	Int(string, int) Event
	Str(string, string) Event
	Float32(string, float32) Event
	Float64(string, float64) Event
	Time(string, time.Time) Event
	Logger() Logger
}

type Logger interface {
	LevelLogger
	With() Event

	WithLevel(Level) Logger
	WithOutput(io.Writer) Logger
	HasLevel(Level) bool
	Level() Level
}

type emptyLogger struct {
}

func (l emptyLogger) WithLevel(Level) Logger      { panic(panicMessage) }
func (l emptyLogger) WithOutput(io.Writer) Logger { panic(panicMessage) }
func (l emptyLogger) With() Event                 { panic(panicMessage) }
func (l emptyLogger) HasLevel(Level) bool         { panic(panicMessage) }
func (l emptyLogger) Trace(msg string)            { panic(panicMessage) }
func (l emptyLogger) Debug(msg string)            { panic(panicMessage) }
func (l emptyLogger) Info(msg string)             { panic(panicMessage) }
func (l emptyLogger) Warn(msg string)             { panic(panicMessage) }
func (l emptyLogger) Error(msg string)            { panic(panicMessage) }
func (l emptyLogger) Fatal(msg string)            { panic(panicMessage) }
func (l emptyLogger) Panic(msg string)            { panic(panicMessage) }
func (l emptyLogger) Write([]byte) (int, error)   { panic(panicMessage) }
func (l emptyLogger) Level() Level                { panic(panicMessage) }

func SetDefaultLogger(l Logger) {
	defaultLogger = l
}

func Str(key, value string) Event {
	return defaultLogger.With().Str(key, value)
}

func Int(key string, value int) Event {
	return defaultLogger.With().Int(key, value)
}

func Trace(msg string) {
	defaultLogger.Trace(msg)
}

func Debug(msg string) {
	defaultLogger.Debug(msg)
}

func Info(msg string) {
	defaultLogger.Info(msg)
}

func Warn(msg string) {
	defaultLogger.Warn(msg)
}

func Error(msg string) {
	defaultLogger.Error(msg)
}

func Fatal(msg string) {
	defaultLogger.Fatal(msg)
}

func Panic(msg string) {
	defaultLogger.Panic(msg)
}

func With() Event {
	return defaultLogger.With()
}

func SetLevel(level Level) {
	defaultLogger = WithLevel(level)
}

func WithLevel(level Level) Logger {
	return defaultLogger.WithLevel(level)
}

func HasLevel(level Level) bool {
	return defaultLogger.HasLevel(level)
}

func HasTrace() bool {
	return HasLevel(LevelTrace)
}

func HasDebug() bool {
	return HasLevel(LevelDebug)
}

func HasInfo() bool {
	return HasLevel(LevelInfo)
}

func HasError() bool {
	return HasLevel(LevelError)
}
