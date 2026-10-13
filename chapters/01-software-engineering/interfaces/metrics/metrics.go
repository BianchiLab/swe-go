package metrics

import (
	"fmt"

		"github.com/BianchiLab/swe-go/chapters/01-software-engineering/interfaces/kv"
		
)

// Sink is implemented by objects that metrics can be reported to.
type Sink interface {
	Put(k string, v interface{}) error
}

// Collector records and publishes metrics.
type Collector struct {
	values map[string]interface{}
}

// NewCollector creates a collector.
func NewCollector() *Collector {
	return &Collector{
		values: make(map[string]interface{}),
	}
}

// Observe records a value for a particular metric.
func (c *Collector) Observe(metric string, value interface{}) {
	c.values[metric] = value
}

// ReportMetrics writes the collected metrics to a Sink.
func (c *Collector) ReportMetrics(s Sink) error {
	for k, v := range c.values {
		if err := s.Put(k, v); err != nil {
			return fmt.Errorf("write metric %q: %w", k, err)
		}
	}
	return nil
}

// Ensure the real KV store implements Sink.
var _ Sink = (*kv.Store)(nil)