package exporter_test

import (
	"testing"

	"github.com/mnadeem/volume_exporter/exporter"
	"github.com/prometheus/client_golang/prometheus"
)

func TestDesc(t *testing.T) {
	volOpts := exporter.VolumeOpts{}
	volOpts.Options = append(volOpts.Options, exporter.VolumeOpt{Name: "name", Path: "path"})

	collector := exporter.NewVolumeCollector(&volOpts)

	desc := make(chan *prometheus.Desc)

	go func() {
		defer close(desc)
		collector.Describe(desc)

	}()

	for i := 1; i <= 4; i++ {
		m := <-desc
		if m == nil {
			t.Error("expected metric but got nil")
		}
		t.Logf("%s", m.String())
	}
	extraMetrics := 0
	for <-desc != nil {
		extraMetrics++
	}
	if extraMetrics > 0 {
		t.Errorf("expected closed channel, got %d extra metrics", extraMetrics)
	}
}
