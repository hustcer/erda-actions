FROM --platform=$TARGETPLATFORM registry.erda.cloud/erda/terminus-golang:1.17.12 as builder

MAINTAINER shenli shenli@terminus.io

COPY . /go/src/github.com/erda-project/erda-actions
WORKDIR /go/src/github.com/erda-project/erda-actions

ARG ARCH

# go build
RUN GOOS=linux GOARCH=$ARCH go build -o /assets/run /go/src/github.com/erda-project/erda-actions/actions/ios/1.0/internal/cmd/main.go

FROM --platform=$TARGETPLATFORM registry.erda.cloud/erda/pipeline-resource:base
COPY --from=builder /assets /opt/action