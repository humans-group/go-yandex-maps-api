package client

import (
	"log"
	"io"
)

// ErrLogger is an implementation of StdLogger that geo uses to log its error messages.
var ErrLogger StdLogger = log.New(io.Discard, "[Geo][Err]", log.LstdFlags)

// DebugLogger is an implementation of StdLogger that geo uses to log its debug messages.
var DebugLogger StdLogger = log.New(io.Discard, "[Geo][Debug]", log.LstdFlags)

// StdLogger is a interface for logging libraries.
type StdLogger interface {
	Printf(string, ...interface{})
}