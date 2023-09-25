FROM        busybox
LABEL maintainer="Mohammad Nadeem<coolmind182006@gmail.com>"

COPY ./volume_exporter /bin/volume_exporter

USER 1001

ENTRYPOINT [ "/bin/volume_exporter" ]
CMD        [ "--volume-dir=bin:/bin", \
             "--web.listen-address=:9888" ]
