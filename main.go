package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/mnadeem/volume_exporter/exporter"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/prometheus/common/log"
	"github.com/prometheus/common/version"
)

type volumeNamePathFlag []string

func (v *volumeNamePathFlag) Set(value string) error {
	*v = append(*v, value)
	return nil
}

func (v *volumeNamePathFlag) String() string {
	return fmt.Sprint(*v)
}

type volumeFlags struct {
	volumeNamePath volumeNamePathFlag
}

func main() {

	var (
		listenAddress = flag.String("web.listen-address", ":9888", "Address to listen on for web interface and telemetry.")
		metricPath    = flag.String("web.telemetry-path", "/metrics", "Path under which to expose metrics.")

		volFlags = volumeFlags{}
	)

	log.Infoln("Starting volume_exporter")

	flag.Var(&volFlags.volumeNamePath, "volume-dir", "Volumes to report, the format is volumeName:VolumeDir;\n For example ==> logs:/app/logs; can be used multiple times to provide more than one value")
	flag.Parse()

	if len(volFlags.volumeNamePath) < 1 {
		log.Infoln("Missing volume-dir")
		log.Infoln()
		flag.Usage()
		os.Exit(1)
	}

	volOpts := exporter.VolumeOpts{}

	for _, np := range volFlags.volumeNamePath {
		name, path := splitFlag(np)
		log.Infof("Directory Name : %s, Path : %s", name, path)
		volOpts.Options = append(volOpts.Options, exporter.VolumeOpt{Name: name, Path: path})
	}

	exporter.Register(&volOpts)

	log.Infoln("Starting volume_exporter", version.Info())
	log.Infoln("Build context", version.BuildContext())

	log.Fatal(serverMetrics(*listenAddress, *metricPath))
}

// Split colon separated flag into two fields
func splitFlag(s string) (string, string) {

	if len(s) == 0 {
		log.Fatalln("Nothing to Monitor")
		os.Exit(1)
	}

	slice := strings.SplitN(s, ":", 2)

	if len(slice) == 1 {
		log.Fatalf("Invalid option %s", s)
		os.Exit(1)
	}
	name := strings.TrimSpace(slice[0])
	path := strings.TrimSpace(slice[1])

	if len(name) == 0 {
		log.Fatalf("Invalid Name on %s", s)
		os.Exit(1)
	}

	exists, err := exists(path)
	if err != nil {
		log.Fatalf("Error validating if path exists %s , error %v", path, err)
		os.Exit(1)
	}
	if !exists {
		log.Fatalf("Invalid Path %s", path)
		os.Exit(1)
	}

	return name, path
}

func exists(path string) (bool, error) {
	stat, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if stat.IsDir() {
		return true, nil
	}
	return false, err
}

func serverMetrics(listenAddress, metricsPath string) error {
	http.Handle(metricsPath, promhttp.Handler())
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`
			<html>
			<head><title>Volume Exporter Metrics</title></head>
			<body>
			<h1>Volume Exporter Metrics</h1>
			<p><a href='` + metricsPath + `'>Metrics</a></p>
			</body>
			</html>
		`))
	})

	log.Infof("Starting Server: %s", listenAddress)
	return http.ListenAndServe(listenAddress, nil)
}
