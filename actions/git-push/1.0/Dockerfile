FROM --platform=$TARGETPLATFORM registry.erda.cloud/erda/terminus-golang:1.17.12 AS builder

COPY . /go/src/github.com/erda-project/erda-actions
WORKDIR /go/src/github.com/erda-project/erda-actions

ARG ARCH

# go build
RUN GOOS=linux GOARCH=$ARCH go build -o /opt/action/run github.com/erda-project/erda-actions/actions/git-push/1.0/internal/cmd
COPY actions/git-push/1.0/internal/git-push.sh /opt/action/git-push.sh

FROM --platform=$TARGETPLATFORM registry.erda.cloud/erda/terminus-alpine:base
COPY --from=builder /opt/action/run /opt/action/run
COPY --from=builder /opt/action/git-push.sh /opt/action/
