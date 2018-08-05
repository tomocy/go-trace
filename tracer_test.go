package trace

import (
	"bytes"
	"testing"
)

func TestNew(t *testing.T) {
	var buf bytes.Buffer
	tracer := New(&buf)
	if tracer == nil {
		t.Error("tracer is nil")
		return
	}

	want := "hello, tracer\n"
	tracer.Trace(want)
	have := buf.String()
	if have != want {
		t.Errorf("have %v, but want %v", have, want)
	}
}
