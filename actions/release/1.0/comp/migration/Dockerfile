# Build app image
FROM --platform=$TARGETPLATFORM registry.erda.cloud/erda-actions/migration:1.0.2

RUN echo "Asia/Shanghai" | tee /etc/timezone

COPY sql/ /tmp/db/migration
RUN chmod +x /entrypoint.sh

WORKDIR /

CMD ["/entrypoint.sh"]