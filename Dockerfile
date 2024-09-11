FROM gcr.io/distroless/static-debian11:nonroot
ENTRYPOINT ["/baton-ivanti"]
COPY baton-ivanti /