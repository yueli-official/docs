FROM golang:1.25.12-bookworm AS build

WORKDIR /src
COPY go.mod go.sum ./
RUN go mod download

COPY api ./api
COPY cmd ./cmd
COPY internal ./internal
RUN CGO_ENABLED=0 GOOS=linux go build -trimpath -ldflags="-s -w" -o /out/docs-api ./cmd/docs
RUN CGO_ENABLED=0 GOOS=linux go build -trimpath -ldflags="-s -w" -o /out/docs-healthcheck ./cmd/healthcheck
RUN CGO_ENABLED=0 GOOS=linux go build -trimpath -ldflags="-s -w" -o /out/docs-bindingcheck ./cmd/bindingcheck
RUN CGO_ENABLED=0 GOOS=linux go build -trimpath -ldflags="-s -w" -o /out/docs-bootstrap ./cmd/bootstrap

FROM migrate/migrate:v4.19.1 AS migration
COPY manifest/sql/migrations /migrations

FROM gcr.io/distroless/static-debian12:nonroot AS bindingcheck
COPY --from=build /out/docs-bindingcheck /usr/local/bin/docs-bindingcheck
USER nonroot:nonroot
ENTRYPOINT ["/usr/local/bin/docs-bindingcheck"]

FROM gcr.io/distroless/static-debian12:nonroot AS bootstrap
COPY --from=build /out/docs-bootstrap /usr/local/bin/docs-bootstrap
USER nonroot:nonroot
ENTRYPOINT ["/usr/local/bin/docs-bootstrap"]

FROM gcr.io/distroless/static-debian12:nonroot AS runtime
WORKDIR /app
COPY --from=build /out/docs-api /usr/local/bin/docs-api
COPY --from=build /out/docs-healthcheck /usr/local/bin/docs-healthcheck
COPY --chown=nonroot:nonroot manifest ./manifest

ENV GF_SERVER_ADDRESS=0.0.0.0:8086
ENV OTEL_TRACES_EXPORTER=none

EXPOSE 8086
USER nonroot:nonroot
ENTRYPOINT ["/usr/local/bin/docs-api"]
