package tracer

type Tracer interface {
	Trace(...interface{})
}
