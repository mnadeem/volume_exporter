## Volume Exporter

[![Build Status](https://travis-ci.com/mnadeem/volume_exporter.svg?branch=master)](https://travis-ci.com/mnadeem/volume_exporter)
[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://opensource.org/licenses/Apache-2.0)
[![Docker](https://img.shields.io/badge/docker-master-brightgreen.svg)](https://hub.docker.com/repository/docker/mnadeem/volume_exporter) 
[![Docker Pulls](https://img.shields.io/docker/pulls/mnadeem/volume_exporter.svg)](https://hub.docker.com/r/mnadeem/volume_exporter)
[![Go Report Card](https://goreportcard.com/badge/github.com/mnadeem/volume_exporter)](https://goreportcard.com/report/github.com/mnadeem/volume_exporter)


Useful to monitor disk/volume/PV storage, for various reasons

* Host path cannot to be mounted to container/Deamonset
* [In ability](https://bugzilla.redhat.com/show_bug.cgi?id=1373288) of Cloud provider to gather information about PV


This is what happens in corporates, multiple projects run on same and/or different nodes, and hence host path mount cannot be granted, further node scrapping cannot be done as well, since it requires cluster role, as in a cluster multiple projects/namespace runs, and hence cluster role is not granted to project teams.

Above I have listed some of the use cases ( real world ) where in node exporter cannot be deployed, should this stop us from monitoring pv/volumes/disk?

**This is altogether a different exporter and does not duplicate any existing exporters. Its area of focus (Where ever node exporter cannot be deployed) is altogether different.**

**It just fills the vacuum**

## Wiki

Refer project [wiki](https://github.com/mnadeem/volume_exporter/wiki) for more details

## Running

### Command Line locally


```bash 
go run main.go --volume-dir=practices:E:\practices
```

#### Usage

```
Usage of Temp\go-build869878202\b001\exe\main.exe:
  -volume-dir value
        Volumes to report, the format is volumeName:VolumeDir;
         For example ==> logs:/app/logs; can be used multiple times to provide more than one value
  -web.listen-address string
        Address to listen on for web interface and telemetry. (default ":9888")
  -web.telemetry-path string
        Path under which to expose metrics. (default "/metrics")
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

## Config

|Flag |	Description|
| ---------------------------- | -------------------------------------------- | 
| web.listen-address |	Address to listen on for web interface and telemetry. Default is 9888|
| web.telemetry-path |	Path under which to expose metrics. Default is /metrics|
| volume-dir	 | volumes to report, the format is volumeName:VolumeDir, For example ==> logs:/app/logs, you can use this flag multiple times to provide multiple volumes|


## Exporterd Metrics

| metrics	| Type |	Description |
| --------------------------------------------------------- | ----------- |  ------------------------------------- |
| volume_bytes_total{volume_name=”someName”, volume_path=”/some/path”} |	Gauge	| Total size of the volume/disk | 
| volume_bytes_free{volume_name=”someName”, volume_path=”/some/path”}	| Gauge	| Free size of the volume/disk | 
| volume_bytes_used{volume_name=”someName”, volume_path=”/some/path”} |	Gauge |	Used size of volume/disk | 

Here is a sample metrics exporterd by running  

```bash
docker run --rm -p 9888:9888  -it docker.repo1.uhc.com/mnadeem/volume_exporter:latest  -volume-dir=bin:/bin -volume-dir=etc:/etc
```

```
# HELP volume_bytes_free Free size of the volume/disk
# TYPE volume_bytes_free gauge
volume_bytes_free{volume_name="bin",volume_path="/bin"} 4.341569536e+10
volume_bytes_free{volume_name="etc",volume_path="/etc"} 4.341569536e+10
# HELP volume_bytes_total Total size of the volume/disk
# TYPE volume_bytes_total gauge
volume_bytes_total{volume_name="bin",volume_path="/bin"} 6.391887872e+10
volume_bytes_total{volume_name="etc",volume_path="/etc"} 6.391887872e+10
# HELP volume_bytes_used Used size of volume/disk
# TYPE volume_bytes_used gauge
volume_bytes_used{volume_name="bin",volume_path="/bin"} 2.050318336e+10
volume_bytes_used{volume_name="etc",volume_path="/etc"} 2.050318336e+10
# HELP volume_exporter_build_info A metric with a constant '1' value labeled by version, revision, branch, and goversion from which volume_exporter was built.
# TYPE volume_exporter_build_info gauge
volume_exporter_build_info{branch="",goversion="go1.15",revision="",version=""} 1
# HELP volume_percentage_used Percentage of volume/disk Utilization
# TYPE volume_percentage_used gauge
volume_percentage_used{volume_name="bin",volume_path="/bin"} 32.07688208958619
volume_percentage_used{volume_name="etc",volume_path="/etc"} 32.07688208958619
```

## Why Volume Exporter

* Because there are cases where in you dont have permission to mounth host path
* Beacause there are cases where you dont have access to nodes.
* Very light weight (just 6.84 MB of image size)
* Configurable, you can pass multiple volumes to track.


## Support
If you need help using volume_exporter feel free to drop an email or [create an issue](https://github.com/mnadeem/volume_exporter/issues/new)  (**preferred**)

## Contributions
To help development you are encouraged to  
* Provide suggestion/feedback/Issue
* pull requests for new features
* Star :star2: the project


[![View My profile on LinkedIn](https://static.licdn.com/scds/common/u/img/webpromo/btn_viewmy_160x33.png)](https://in.linkedin.com/pub/nadeem-mohammad/17/411/21)

## Inspired From

* [Minio Disk](https://github.com/minio/minio/blob/master/pkg/disk/disk.go)
* [configmap-reload](https://github.com/jimmidyson/configmap-reload)
* [kafka_exporter](https://github.com/danielqsj/kafka_exporter)