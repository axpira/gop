package log

import (
	"fmt"
	"time"
)

// LogMarshaler provides a way to marshal the object using FieldBuilder
// it's used by Marshal or Any methods
type LogMarshaler interface {
	MarshalLog(FieldBuilder)
}

// FieldBuilder represents fields in a log context.
// Each method adds a key with the value as type
type FieldBuilder interface {
	LoggerOption

	// Str adds the field key with a value as a string to the FieldBuilder
	Str(string, string) FieldBuilder
	// Bool adds the field key with a value as a bool to the FieldBuilder
	Bool(string, bool) FieldBuilder
	// Bytes adds the field key with a value as a slice of byte to the FieldBuilder
	Bytes(string, []byte) FieldBuilder

	// Float32 adds the field key with a value as a float32 to the FieldBuilder
	Float32(string, float32) FieldBuilder
	// Float64 adds the field key with a value as a float64 to the FieldBuilder
	Float64(string, float64) FieldBuilder

	// Msg adds the message for a log to the FieldBuilder.
	Msg(string) FieldBuilder
	// Msgf adds the message for a log, using formatter, to the FieldBuilder.
	Msgf(string, ...interface{}) FieldBuilder

	// Interface adds the field key with value as a interface to the FieldBuilder.
	// This method will find the type of value to add in the log
	Interface(string, interface{}) FieldBuilder
	// Marshal adds the field key with value as a interface to the FieldBuilder.
	// This method will try to Marshal as LogMarshaler interface or
	//  will marshal as expected by implementation
	Marshal(string, interface{}) FieldBuilder

	// Complex64 adds the field key with a value as a complex64 to the FieldBuilder
	Complex64(string, complex64) FieldBuilder
	// Complex128 adds the field key with a value as a complex128 to the FieldBuilder
	Complex128(string, complex128) FieldBuilder

	// Fields adds for each keys with values the FieldBuilder.
	Fields(map[string]interface{}) FieldBuilder
	// Dict adds the field key with a value as a FieldBuilder to the FieldBuilder
	Dict(string, FieldBuilder) FieldBuilder

	// Int adds the field key with a value as a int to the FieldBuilder
	Int(string, int) FieldBuilder
	// Int8 adds the field key with a value as a int8 to the FieldBuilder
	Int8(string, int8) FieldBuilder
	// Int16 adds the field key with a value as a int16 to the FieldBuilder
	Int16(string, int16) FieldBuilder
	// Int32 adds the field key with a value as a int32 to the FieldBuilder
	Int32(string, int32) FieldBuilder
	// Int64 adds the field key with a value as a int64 to the FieldBuilder
	Int64(string, int64) FieldBuilder

	// Uint adds the field key with a value as a uint to the FieldBuilder
	Uint(string, uint) FieldBuilder
	// Uint8 adds the field key with a value as a uint8 to the FieldBuilder
	Uint8(string, uint8) FieldBuilder
	// Uint16 adds the field key with a value as a uint16 to the FieldBuilder
	Uint16(string, uint16) FieldBuilder
	// Uint32 adds the field key with a value as a uint32 to the FieldBuilder
	Uint32(string, uint32) FieldBuilder
	// Uint64 adds the field key with a value as a uint64 to the FieldBuilder
	Uint64(string, uint64) FieldBuilder

	// Timef adds the field key with a value as a time formatted with format to the FieldBuilder
	Timef(key string, value time.Time, format string) FieldBuilder
	// Time adds the field key with a value as a time to the FieldBuilder
	Time(string, time.Time) FieldBuilder

	// Dur adds the field key with a value as a duration to the FieldBuilder
	Dur(string, time.Duration) FieldBuilder

	// Stringer adds the field key with a value as a fmt.Stringer to the FieldBuilder
	Stringer(string, fmt.Stringer) FieldBuilder

	// Error adds the field key with a value as a error to the FieldBuilder
	Error(string, error) FieldBuilder
	// Err adds error field with value to the FieldBuilder.
	Err(error) FieldBuilder
}
