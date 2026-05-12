FROM scratch AS fs
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

FROM timoreymann/ubuntu-runtime:26.6
LABEL org.opencontainers.image.title="yal" \
      org.opencontainers.image.description="A simple link hub, to display and search links. Allows easy branding, runs with the least privileges and is simple to use." \
      org.opencontainers.image.ref.name="main" \
      org.opencontainers.image.licenses='GNU GPL v3' \
      org.opencontainers.image.vendor="Timo Reymann <mail@timo-reymann.de>" \
      org.opencontainers.image.authors="Timo Reymann <mail@timo-reymann.de>" \
      org.opencontainers.image.url="https://github.com/timo-reymann/yal" \
      org.opencontainers.image.documentation="https://github.com/timo-reymann/yal" \
      org.opencontainers.image.source="https://github.com/timo-reymann/yal.git"

ENV YAL_PORT=2024 \
    YAL_MASCOT=mascot \
    YAL_LOGO=logo \
    YAL_BACKGROUND=background \
    YAL_FAVICON=favicon \
    YAL_CONFIG_FOLDER=/app/config \
    YAL_IMAGES_FOLDER=/app/images

WORKDIR /app
COPY --from=bin /bin/yal .
COPY --from=fs / ./

ENTRYPOINT [ "/app/yal" ]
CMD [ "-serve" ]
