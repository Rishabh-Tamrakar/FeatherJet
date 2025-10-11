package metrics

import (
	"testing"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/testutil"
)

func TestMetrics(t *testing.T) {
	// Test RequestsTotal counter
	RequestsTotal.WithLabelValues("GET", "/test", "200").Inc()
	
	expected := `
# HELP http_requests_total Total number of HTTP requests
# TYPE http_requests_total counter
http_requests_total{method="GET",path="/test",status="200"} 1
`
	if err := testutil.CollectAndCompare(RequestsTotal, expected); err != nil {
		t.Errorf("unexpected collecting result:\n%s", err)
	}

	// Test RequestDuration histogram
	RequestDuration.WithLabelValues("GET", "/test").Observe(0.1)
	
	// Reset metrics for other tests
	prometheus.Unregister(RequestsTotal)
	prometheus.Unregister(RequestDuration)
}