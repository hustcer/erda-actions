FROM --platform=$TARGETPLATFORM registry.erda.cloud/erda/terminus-golang:1.11.2 AS builder

COPY . /go/src/github.com/erda-project/erda-actions
WORKDIR /go/src/github.com/erda-project/erda-actions

ARG ARCH

# go build
RUN GOOS=linux GOARCH=$ARCH go build -o /assets/run actions/java-lint/1.0/internal/cmd/*.go

FROM --platform=$TARGETPLATFORM registry.erda.cloud/erda/pipeline-resource:base AS action
COPY actions/java-lint/1.0/internal/settings.xml /root/.m2/settings.xml
COPY --from=builder /assets /opt/action

FROM action