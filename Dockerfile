FROM gcr.io/distroless/static-debian12:nonroot
ENTRYPOINT ["/cone"]
COPY cone /