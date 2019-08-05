package traces

import (
	"github.com/opentracing/opentracing-go/mocktracer"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBufferedSpans_Set_OverBuffer_Bounded(t *testing.T) {
	spans, err := NewBufferedSpanStore(1)
	assert.NoError(t, err)

	tracer := mocktracer.New()
	span1 := tracer.StartSpan("span1")
	span2 := tracer.StartSpan("span2")

	spans.Set("span1", span1)
	spans.Set("span2", span2)
	assert.Equal(t, 1, spans.Count())

	// check that the span is the second span
	s2, ok := spans.Get("span2")
	assert.True(t, ok)
	assert.Equal(t, span2, s2)
}

func TestBufferedSpans_Delete(t *testing.T) {
	spans, err := NewBufferedSpanStore(1)
	assert.NoError(t, err)

	tracer := mocktracer.New()
	span1 := tracer.StartSpan("span1")
	spans.Set("span1", span1)
	spans.Delete("span1")
	assert.Equal(t, 0, spans.Count())
	assert.Nil(t, spans.buf[0])
}
