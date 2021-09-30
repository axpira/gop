package log

type Level uint8

const (
	DisabledLevel Level = iota
	TraceLevel
	DebugLevel
	InfoLevel
	WarnLevel
	ErrorLevel
	FatalLevel
	PanicLevel
	NoLevel
)
