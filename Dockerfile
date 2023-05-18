FROM gcr.io/distroless/static-debian11:nonroot
ENTRYPOINT ["/cone"]
COPY cone /