FROM scratch as fs
COPY ./config /config
COPY ./images /images
COPY ./icons /icons

FROM busybox AS bin
COPY ./dist /binaries
RUN if [[ "$(arch)" == "x86_64" ]]; then \
        architecture="amd64"; \
    else \
        architecture="arm64"; \
    fi; \
    cp /binaries/yal_linux-${architecture} /bin/yal && \
    chmod +x /bin/yal && \
    chown 65532:65532 /bin/yal

FROM scratch

LABEL org.opencontainers.image.title="yal"
LABEL org.opencontainers.image.description="A simple link hub, to display and search links. Allows easy branding, runs with the least privileges and is simple to use."
LABEL org.opencontainers.image.ref.name="main"
LABEL org.opencontainers.image.licenses='GNU GPL v3'
LABEL org.opencontainers.image.vendor="Timo Reymann <mail@timo-reymann.de>"
LABEL org.opencontainers.image.authors="Timo Reymann <mail@timo-reymann.de>"
LABEL org.opencontainers.image.url="https://github.com/timo-reymann/yal"
LABEL org.opencontainers.image.documentation="https://github.com/timo-reymann/yal"
LABEL org.opencontainers.image.source="https://github.com/timo-reymann/yal.git"

ENV YAL_PORT=2024
ENV YAL_MASCOT=mascot
ENV YAL_LOGO=logo
ENV YAL_BACKGROUND=background
ENV YAL_FAVICON=favicon
ENV YAL_CONFIG_FOLDER=/app/config
ENV YAL_IMAGES_FOLDER=/app/images

COPY --from=gcr.io/distroless/static-debian12:nonroot / /
USER nonroot

WORKDIR /app
COPY --from=bin /bin/yal .
COPY --from=fs / ./

ENTRYPOINT [ "/app/yal" ]
CMD [ "-serve" ]
