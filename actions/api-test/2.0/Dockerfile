FROM --platform=$TARGETPLATFORM registry.erda.cloud/erda/terminus-golang:1.17.12 AS builder

COPY . /go/src/github.com/erda-project/erda-actions
WORKDIR /go/src/github.com/erda-project/erda-actions

ARG ARCH

# go build
RUN GOOS=linux GOARCH=$ARCH go build -o /assets/run actions/api-test/2.0/internal/*.go

FROM --platform=$TARGETPLATFORM registry.erda.cloud/erda/terminus-centos:base
RUN yum install -y https://dl.fedoraproject.org/pub/epel/epel-release-latest-7.noarch.rpm && yum install -y jq
COPY --from=builder /assets /opt/action
