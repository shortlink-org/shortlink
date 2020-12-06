FROM scratch
COPY api /usr/bin/api
ENTRYPOINT ["/usr/bin/api"]
