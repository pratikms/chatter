package trace

import "io"

// Tracer is the interface that describes an object capable of tracing events throughout code
type Tracer interface {
	Trace(...interface{})
}

// New creates a new Tracer that will write the output to
// the specified io.Writer.
func New(w io.Writer) Tracer {
	return nil
}
