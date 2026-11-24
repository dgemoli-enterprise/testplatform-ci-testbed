FROM docker.io/library/alpine:3.23
COPY hello /hello
USER 65532:65532
ENTRYPOINT ["/hello"]
