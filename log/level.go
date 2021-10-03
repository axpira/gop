package log

type Level uint8

const (
	NoLevel Level = iota
	TraceLevel
	DebugLevel
	InfoLevel
	WarnLevel
	ErrorLevel
	FatalLevel
	PanicLevel
	DisabledLevel
)
