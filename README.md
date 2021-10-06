[![Go Reference](https://pkg.go.dev/badge/github.com/axpira/gop.svg)](https://pkg.go.dev/github.com/axpira/gop)
[![Go Report Card](https://goreportcard.com/badge/github.com/axpira/gop)](https://goreportcard.com/report/github.com/axpira/gop)

# gop

It's just a set of interfaces to use somethings with Golang.

The big advantage you can use this interfaces without need to think the implementation.

# How to use

## LOG

You must choose one implementation, this is some examples:

- "github.com/axpira/goplogjson"
_ "github.com/axpira/goplogadapter/logrus"


And to use:

```go
package main

import (
	"os"

	"github.com/axpira/gop/log"
	_ "github.com/axpira/goplogjson"
	// _ "github.com/axpira/goplogadapter/logrus" // uncommenting this line and commenting the above line
                                                // the Logrus will be use as default implementation
)

func main() {
	log.Info("Hello World")
	// out: {"msg":"Hello World","level":"info","time":"2021-09-29T07:43:34-03:00"}

	// Adding fields
	log.Inf(log.
		Str("str_field", "hello").
		Int("int_field", 42).
		Msg("Hello World"),
	)
	// out: {"str_field":"hello","int_field":42,"msg":"Hello World","level":"info","time":"2021-09-29T07:46:12-03:00"}

	// Creating child log
	l := log.With(log.
		Str("str_field", "hello").
		Int("int_field", 42),
	)
	// out: {"msg":"Hello World","level":"info","str_field":"hello","int_field":42,"time":"2021-09-29T07:49:29-03:00"}

	l.Info("Hello World")
  // out: {"msg":"Hello World","level":"info","str_field":"hello","int_field":42,"time":"2021-09-29T08:00:48-03:00"}
	l.Inf(log.
		Bool("another_field", true).
		Msg("Hello World"),
	)
	// out: {"another_field":true,"msg":"Hello World","level":"info","str_field":"hello","int_field":42,"time":"2021-09-29T07:49:29-03:00"}

}
```
