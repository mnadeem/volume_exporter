package exporter

import (
	"fmt"
	"log"

	"github.com/mnadeem/volume_exporter/disk"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/common/version"
)

const (
	namespace = "volume" //for Prometheus metrics.
)

//VolumeOpts is for options
type VolumeOpts struct {
	Options []VolumeOpt
}

//VolumeOpt is for option
type VolumeOpt struct {
	Name string
	Path string
}

//Define a struct for you collector that contains pointers
//to prometheus descriptors for each metric you wish to expose.
//Note you can also include fields of other types if they provide utility
//but we just won't be exposing them as metrics.
type volumeCollector struct {
	volumeBytesTotal *prometheus.Desc
	volumeBytesFree  *prometheus.Desc
	volumeBytesUsed  *prometheus.Desc

	volOptions VolumeOpts
}

//You must create a constructor for you collector that
//initializes every descriptor and returns a pointer to the collector
func newVolumeCollector(opts *VolumeOpts) *volumeCollector {
	return &volumeCollector{
		volumeBytesTotal: prometheus.NewDesc(prometheus.BuildFQName(namespace, "", "bytes_total"),
			"Total size of the volume/disk",
			[]string{"name", "path"}, nil,
		),
		volumeBytesFree: prometheus.NewDesc(prometheus.BuildFQName(namespace, "", "bytes_free"),
			"Free size of the volume/disk",
			[]string{"name", "path"}, nil,
		),
		volumeBytesUsed: prometheus.NewDesc(prometheus.BuildFQName(namespace, "", "bytes_used"),
			"Used size of volume/disk",
			[]string{"name", "path"}, nil,
		),
		volOptions: *opts,
	}
}

//Each and every collector must implement the Describe function.
//It essentially writes all descriptors to the prometheus desc channel.
func (collector *volumeCollector) Describe(ch chan<- *prometheus.Desc) {

	//Update this section with the each metric you create for a given collector
	ch <- collector.volumeBytesTotal
	ch <- collector.volumeBytesFree
	ch <- collector.volumeBytesUsed
}

//Collect implements required collect function for all promehteus collectors
func (collector *volumeCollector) Collect(ch chan<- prometheus.Metric) {

	//Implement logic here to determine proper metric value to return to prometheus
	//for each descriptor or call other functions that do so.
	for _, opt := range collector.volOptions.Options {

		di, err := disk.GetInfo(opt.Path)
		if err != nil {
			log.Fatal(err)
		}

		percentage := (float64(di.Total-di.Free) / float64(di.Total)) * 100
		fmt.Printf("%d of %d disk space used (%0.2f%%)\n", di.Total-di.Free, di.Total, percentage)

		//Write latest value for each metric in the prometheus metric channel.
		//Note that you can pass CounterValue, GaugeValue, or UntypedValue types here.
		ch <- prometheus.MustNewConstMetric(collector.volumeBytesTotal, prometheus.GaugeValue, float64(di.Total), opt.Name, opt.Path)
		ch <- prometheus.MustNewConstMetric(collector.volumeBytesFree, prometheus.GaugeValue, float64(di.Free), opt.Name, opt.Path)
		ch <- prometheus.MustNewConstMetric(collector.volumeBytesUsed, prometheus.GaugeValue, float64(di.Used), opt.Name, opt.Path)
	}
}

// Register registers the volume metrics
func Register(options *VolumeOpts) {
	collector := newVolumeCollector(options)
	prometheus.MustRegister(collector)
	prometheus.MustRegister(version.NewCollector("volume_exporter"))
	prometheus.Unregister(prometheus.NewGoCollector())
}
