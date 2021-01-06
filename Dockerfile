FROM        busybox
LABEL maintainer="Mohammad Nadeem<coolmind182006@gmail.com>"

COPY ./volume_exporter /bin/volume_exporter

ENTRYPOINT [ "/bin/volume_exporter" ]