FROM busybox

LABEL maintainer="Mohammad Nadeem<coolmind182006@gmail.com>"

RUN adduser -D vExporter
COPY --chown=vExporter ./volume_exporter.* /bin/volume_exporter

USER vExporter

ENTRYPOINT [ "/bin/volume_exporter" ]
CMD        [ "--volume-dir=bin:/bin", \
             "--web.listen-address=:9888" ]
