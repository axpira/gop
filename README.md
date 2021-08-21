# What

It's just a set of interfaces to use somethings with Golang.

The big advantage you can use this interfaces without need to mint the implementation.
For example LOG you can log with this:
```go
log.Info("Hello World")
```
And can use any of implementations to print this log correct at output


# How to use

## Gop Log (default implementation)
```go
package main

import (
	"github.com/axpira/gop/log"
	"github.com/axpira/goplog"
	"github.com/axpira/goplog/enc/console"
)

func main() {
	// Default Level is INFO with JSON format
	log.SetDefaultLogger(goplog.New())
	log.Info("Hello World")
	log.Debug("No Hello Debug")

	// Changing to CONSOLE format
	goplog.SetFormat(console.Console{})
	// Changing default level do DEBUG
	log.SetLevel(log.LevelDebug)
	log.Debug("Hello Debug")

	// Creating child log
	l := log.WithLevel(log.LevelTrace).With().Str("foo", "bar").Int("int", 42).Logger()
	l.With().Str("another", "field").Trace("Hello Trace")
}
```
