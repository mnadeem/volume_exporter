FROM        quay.io/prometheus/busybox:latest
LABEL maintainer="Mohammad Nadeem<coolmind182006@gmail.com>"

COPY volume_exporter /bin/volume_exporter

EXPOSE     9308
ENTRYPOINT [ "/bin/volume_exporter" ]