package trace

import (
	"fmt"
	"io"
)

// Tracer is the interface that describes an object capable of tracing events throughout code
type Tracer interface {
	Trace(...interface{})
}

type tracer struct {
	out io.Writer
}

type nilTracer struct{}

func (t *nilTracer) Trace(a ...interface{}) {}

func (t *tracer) Trace(a ...interface{}) {
	fmt.Fprint(t.out, a...)
	fmt.Fprintln(t.out)
}

// New creates a new Tracer that will write the output to
// the specified io.Writer.
func New(w io.Writer) Tracer {
	return &tracer{out: w}
}

func Off() Tracer {
	return &nilTracer{}
}
