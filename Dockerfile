FROM scratch
COPY server /
COPY secrets.yaml /
ENTRYPOINT ["/server"]
