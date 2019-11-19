package eventsources

import (
	"github.com/opentracing/opentracing-go"
	"net/http"
)

type SpanState string

type EventState string

const (
	TracePrefix string = "vstrace"
)

const (
	StartState        SpanState = "start"
	EndState          SpanState = "end"
	IntermediaryState SpanState = "intermediary"
	TransitionState   SpanState = "transition"
	UnknownState      SpanState = "unknown"
)

type Event interface {
	SpanID() (string, error)
	OperationName() string
	ParentSpanID() (*string, error)
	IsError() (bool, error)
	State(prev *EventState) (SpanState, error)
	Tags() (map[string]interface{}, error)
}

type EventSource interface {
	Name() string
	ValidatePayload(r *http.Request, secretKey []byte) ([]byte, error)
	Event(r *http.Request, payload []byte) (Event, error)
	Tracer() opentracing.Tracer
}