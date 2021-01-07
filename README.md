## Volume Exporter

[![Build Status](https://travis-ci.com/mnadeem/volume_exporter.svg?branch=master)](https://travis-ci.com/mnadeem/volume_exporter)
[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://opensource.org/licenses/Apache-2.0)

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
docker run  -p 9889:9888 -it mnadeem/volume_exporter --volume-dir=bin:/bin
```
### Deploy It in Cloud
Add as a sidecar

```bash 
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


## Inspired From

* [Minio Disk](https://github.com/minio/minio/blob/master/pkg/disk/disk.go)
* [configmap-reload](https://github.com/jimmidyson/configmap-reload)
* [kafka_exporter](https://github.com/danielqsj/kafka_exporter)