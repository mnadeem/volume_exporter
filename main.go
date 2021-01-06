package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	humanize "github.com/dustin/go-humanize"
	"github.com/mnadeem/volume_exporter/disk"
	_ "github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	listenAddress = flag.String("web.listen-address", ":9888", "Address to listen on for web interface.")
	metricPath    = flag.String("web.metrics-path", "/metrics", "Path under which to expose metrics.")
	volumeDirs    volumeDirsFlag
)

type volumeDirsFlag []string

func (v *volumeDirsFlag) Set(value string) error {
	*v = append(*v, value)
	return nil
}

func (v *volumeDirsFlag) String() string {
	return fmt.Sprint(*v)
}

func main() {

	log.Println("Starting volume_exporter")

	flag.Var(&volumeDirs, "volume-dir", "the config map volume directory to watch for updates; may be used multiple times")
	flag.Parse()

	for _, d := range volumeDirs {
		log.Printf("Watching directory: %q", d)
	}

	di, err := disk.GetInfo("c:\\")
	if err != nil {
		log.Fatal(err)
	}
	percentage := (float64(di.Total-di.Free) / float64(di.Total)) * 100
	fmt.Printf("%s of %s disk space used (%0.2f%%)\n",
		humanize.Bytes(di.Total-di.Free),
		humanize.Bytes(di.Total),
		percentage,
	)

	log.Fatal(serverMetrics(*listenAddress, *metricPath))
}

func serverMetrics(listenAddress, metricsPath string) error {
	http.Handle(metricsPath, promhttp.Handler())
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`
			<html>
			<head><title>Volume Exporter Metrics</title></head>
			<body>
			<h1>ConfigMap Reload</h1>
			<p><a href='` + metricsPath + `'>Metrics</a></p>
			</body>
			</html>
		`))
	})
	log.Printf("Starting Server: %s", listenAddress)
	return http.ListenAndServe(listenAddress, nil)
}
