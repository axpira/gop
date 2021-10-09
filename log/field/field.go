package field

import (
	"time"

	"github.com/axpira/gop/log"
)

// Int starts a new log.FieldBuilder using DefaultLogger,
// with the field key with value as a int to log.FieldBuilder
func Int(key string, value int) log.FieldBuilder {
	return log.DefaultLogger.NewFieldBuilder().Int(key, value)
}

// Fields starts a new log.FieldBuilder using DefaultLogger,
// with adds for each keys with values the log.FieldBuilder.
func Fields(m map[string]interface{}) log.FieldBuilder {
	return log.DefaultLogger.NewFieldBuilder().Fields(m)
}

// Float32 starts a new log.FieldBuilder using DefaultLogger,
// with the field key with value as a float32 to log.FieldBuilder
func Float32(key string, value float32) log.FieldBuilder {
	return log.DefaultLogger.NewFieldBuilder().Float32(key, value)
}

// Float64 starts a new log.FieldBuilder using DefaultLogger,
// with the field key with value as a float64 to log.FieldBuilder
func Float64(key string, value float64) log.FieldBuilder {
	return log.DefaultLogger.NewFieldBuilder().Float64(key, value)
}

// Time starts a new log.FieldBuilder using DefaultLogger,
// with the field key with value as a time to log.FieldBuilder
func Time(key string, value time.Time) log.FieldBuilder {
	return log.DefaultLogger.NewFieldBuilder().Time(key, value)
}

// Dur starts a new log.FieldBuilder using DefaultLogger,
// with the field key with value as a duration to log.FieldBuilder
func Dur(key string, value time.Duration) log.FieldBuilder {
	return log.DefaultLogger.NewFieldBuilder().Dur(key, value)
}

// Err starts a new log.FieldBuilder using DefaultLogger,
// with the field err in log.FieldBuilder
func Err(err error) log.FieldBuilder {
	return log.DefaultLogger.NewFieldBuilder().Err(err)
}

// Error starts a new log.FieldBuilder using DefaultLogger,
// with the field key with value as a error to log.FieldBuilder
func Error(key string, err error) log.FieldBuilder {
	return log.DefaultLogger.NewFieldBuilder().Error(key, err)
}

// Msg starts a new log.FieldBuilder using DefaultLogger,
// with the msg as a message to log.FieldBuilder
func Msg(msg string) log.FieldBuilder {
	return log.DefaultLogger.NewFieldBuilder().Msg(msg)
}

// Msgs starts a new log.FieldBuilder using DefaultLogger,
// with the formatted message with format and args
func Msgf(format string, args ...interface{}) log.FieldBuilder {
	return log.DefaultLogger.NewFieldBuilder().Msgf(format, args...)
}
