FROM busybox

LABEL maintainer="Mohammad Nadeem<coolmind182006@gmail.com>"

RUN adduser -D vExporter
COPY ./volume_exporter.* /bin/volume_exporter

RUN chown -R vExporter:0 /bin/volume_exporter && \
    chmod -R ug+rwx /bin/volume_exporter

USER vExporter

ENTRYPOINT [ "/bin/volume_exporter" ]
CMD        [ "--volume-dir=bin:/bin", \
             "--web.listen-address=:9888" ]
