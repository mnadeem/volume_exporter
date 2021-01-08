package exporter_test

import (
	"io/ioutil"
	"os"
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
			t.Error("expected desc but got nil")
		}
	}
	extraDesc := 0
	for <-desc != nil {
		extraDesc++
	}
	if extraDesc > 0 {
		t.Errorf("expected closed channel, got %d extra desc", extraDesc)
	}
}

func TestCollect(t *testing.T) {
	path, err := ioutil.TempDir(os.TempDir(), "disk-")
	defer os.RemoveAll(path)
	if err != nil {
		t.Fatal(err)
	}

	volOpts := exporter.VolumeOpts{}
	volOpts.Options = append(volOpts.Options, exporter.VolumeOpt{Name: "temp", Path: path})

	collector := exporter.NewVolumeCollector(&volOpts)

	ch := make(chan prometheus.Metric)

	go func() {
		defer close(ch)
		collector.Collect(ch)
	}()

	for i := 1; i <= 4; i++ {
		m := <-ch
		if m == nil {
			t.Error("expected metric but got nil")
		}
	}
	extraMetrics := 0
	for <-ch != nil {
		extraMetrics++
	}
	if extraMetrics > 0 {
		t.Errorf("expected closed channel, got %d extra metrics", extraMetrics)
	}
}
