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

	s := "hello, tracer"
	want := s + "\n"
	tracer.Trace(s)
	have := buf.String()
	if have != want {
		t.Errorf("have %s, but want %s", have, want)
	}
}
