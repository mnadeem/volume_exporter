FROM busybox

LABEL maintainer="Mohammad Nadeem<coolmind182006@gmail.com>"

RUN adduser -D vExporter
COPY ./volume_exporter.* /bin/volume_exporter

RUN chown -R vExporter:0 /bin && \
    chmod -R ug+rwx /bin

USER vExporter

ENTRYPOINT [ "/bin/volume_exporter" ]
CMD        [ "--volume-dir=bin:/bin", \
             "--web.listen-address=:9888" ]
