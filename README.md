## Volume Exporter

Useful to monitor disk/volume/PV storage, for various reasons

* Host path cannot to be mounted to container/Deamonset
* [In ability](https://bugzilla.redhat.com/show_bug.cgi?id=1373288) of Cloud provider to gather information about PV

## Inspired From

* [Minio Disk](https://github.com/minio/minio/blob/master/pkg/disk/disk.go)
* [configmap-reload](https://github.com/jimmidyson/configmap-reload)
* [kafka_exporter](https://github.com/danielqsj/kafka_exporter)