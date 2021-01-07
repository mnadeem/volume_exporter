## Volume Exporter

[![Build Status](https://travis-ci.com/mnadeem/volume_exporter.svg?branch=master)](https://travis-ci.com/mnadeem/volume_exporter)
[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://opensource.org/licenses/Apache-2.0)
[![Docker](https://img.shields.io/badge/docker-master-brightgreen.svg)](https://hub.docker.com/repository/docker/mnadeem/volume_exporter) 
[![Docker Pulls](https://img.shields.io/docker/pulls/mnadeem/volume_exporter.svg)](https://hub.docker.com/r/mnadeem/volume_exporter)
[![Go Report Card](https://goreportcard.com/badge/github.com/mnadeem/volume_exporter)](https://goreportcard.com/report/github.com/mnadeem/volume_exporter)


Useful to monitor disk/volume/PV storage, for various reasons

* Host path cannot to be mounted to container/Deamonset
* [In ability](https://bugzilla.redhat.com/show_bug.cgi?id=1373288) of Cloud provider to gather information about PV

## Wiki

Refer project [wiki](https://github.com/mnadeem/volume_exporter/wiki) for momre details

## Running

### Command Line locally

```bash 
go run main.go --volume-dir=practices:E:\practices
```
### Docker Locally

```bash 
docker run --rm -p 9889:9888 -it mnadeem/volume_exporter --volume-dir=bin:/bin
```
### Deploy It in Cloud
Add as a sidecar

```yaml 
        - name: volume-exporter
          image:  mnadeem/volume_exporter
          imagePullPolicy: "Always"
          args:
            - --volume-dir=prometheus:/prometheus
          ports:
          - name: metrics-volume
            containerPort: 9888
          volumeMounts:
          - mountPath: /prometheus
            name: prometheus-data
            readOnly: true
```

## Exporterd Metrics

Here is a sample metrics exporterd by running  `docker run --rm -p 9888:9888  -it docker.repo1.uhc.com/mnadeem/volume_exporter:latest  -volume-dir=bin:/bin -volume-dir=etc:/etc`


```
# HELP volume_bytes_free Free size of the volume/disk
# TYPE volume_bytes_free gauge
volume_bytes_free{name="bin",path="/bin"} 4.3428974592e+10
volume_bytes_free{name="etc",path="/etc"} 4.3428974592e+10
# HELP volume_bytes_total Total size of the volume/disk
# TYPE volume_bytes_total gauge
volume_bytes_total{name="bin",path="/bin"} 6.391887872e+10
volume_bytes_total{name="etc",path="/etc"} 6.391887872e+10
# HELP volume_bytes_used Used size of volume/disk
# TYPE volume_bytes_used gauge
volume_bytes_used{name="bin",path="/bin"} 2.0489904128e+10
volume_bytes_used{name="etc",path="/etc"} 2.0489904128e+10
# HELP volume_exporter_build_info A metric with a constant '1' value labeled by version, revision, branch, and goversion from which volume_exporter was built.
# TYPE volume_exporter_build_info gauge
volume_exporter_build_info{branch="",goversion="go1.12",revision="",version=""} 1
# HELP volume_percentage_used Percentage of volume/disk Utilization
# TYPE volume_percentage_used gauge
volume_percentage_used{name="bin",path="/bin"} 32.056106956689746
volume_percentage_used{name="etc",path="/etc"} 32.056106956689746
```

## Inspired From

* [Minio Disk](https://github.com/minio/minio/blob/master/pkg/disk/disk.go)
* [configmap-reload](https://github.com/jimmidyson/configmap-reload)
* [kafka_exporter](https://github.com/danielqsj/kafka_exporter)