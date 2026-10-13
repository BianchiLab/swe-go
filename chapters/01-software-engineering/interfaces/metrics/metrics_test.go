package metrics

import (
	"errors"
	"testing"
)

type inMemorySink struct {
	values map[string]any
	err    error
}

// Compile-time check that inMemorySink implements Sink.
var _ Sink = (*inMemorySink)(nil)

func (s *inMemorySink) Put(key string, value any) error {
	if s.err != nil {
		return s.err
	}

	if s.values == nil {
		s.values = make(map[string]any)
	}

	s.values[key] = value
	return nil
}

func TestReportMetrics(t *testing.T) {
	sink := &inMemorySink{}
	collector := NewCollector()

	collector.Observe("requests", 42)
	collector.Observe("latency_ms", 15.5)

	if err := collector.ReportMetrics(sink); err != nil {
		t.Fatalf("ReportMetrics() error = %v", err)
	}

	if got := sink.values["requests"]; got != 42 {
		t.Errorf("requests = %v, want 42", got)
	}

	if got := sink.values["latency_ms"]; got != 15.5 {
		t.Errorf("latency_ms = %v, want 15.5", got)
	}
}

func TestReportMetricsReturnsSinkError(t *testing.T) {
	wantErr := errors.New("sink unavailable")
	sink := &inMemorySink{err: wantErr}
	collector := NewCollector()

	collector.Observe("requests", 42)

	err := collector.ReportMetrics(sink)
	if !errors.Is(err, wantErr) {
		t.Errorf("ReportMetrics() error = %v, want %v", err, wantErr)
	}
}
