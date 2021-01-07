FROM        busybox
LABEL maintainer="Mohammad Nadeem<coolmind182006@gmail.com>"

COPY ./main /bin/volume_exporter

ENTRYPOINT [ "/bin/volume_exporter" ]