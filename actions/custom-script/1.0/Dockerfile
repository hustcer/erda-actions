FROM --platform=$TARGETPLATFORM registry.erda.cloud/retag/buildkit:v0.11.3 as buildkit
FROM --platform=$TARGETPLATFORM registry.erda.cloud/retag/golang:1.19.7-alpine3.16

ENV GOPATH /go
ENV PATH $GOPATH/bin:/usr/local/go/bin:$PATH
ENV GOPROXY=https://goproxy.cn/
RUN apk add --no-cache librdkafka-dev bash libc6-compat  \
    openjdk8 nodejs npm yarn docker-cli libgit2-dev git \
    musl-dev gcc curl python3 py-pip openssh pkgconfig \
    make tzdata musl-dev build-base binutils-gold

COPY --from=buildkit /usr/bin/buildctl /usr/bin/buildctl

RUN go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.27.0

RUN mkdir -p /usr/share/maven /usr/share/maven/ref \
  && curl -fsSL -o /tmp/apache-maven.tar.gz https://mirrors.bfsu.edu.cn/apache/maven/maven-3/3.6.3/binaries/apache-maven-3.6.3-bin.tar.gz \
  && tar -xzf /tmp/apache-maven.tar.gz -C /usr/share/maven --strip-components=1 \
  && rm -f /tmp/apache-maven.tar.gz \
  && ln -s /usr/share/maven/bin/mvn /usr/bin/mvn

ARG TARGETARCH

ADD actions/custom-script/1.0/assets/settings.xml /root/.m2/settings.xml

RUN ln -sf /bin/bash /bin/sh

RUN wget http://gosspublic.alicdn.com/ossutil/1.5.0/ossutil64 -O /usr/local/bin/ossutil && chmod 755 /usr/local/bin/ossutil

RUN wget https://github.com/kubernetes-sigs/cri-tools/releases/download/v1.24.1/crictl-v1.24.1-linux-${TARGETARCH}.tar.gz \
    && tar zxvf crictl-v1.24.1-linux-${TARGETARCH}.tar.gz -C /usr/local/bin \
    && rm -f crictl-v1.24.1-linux-${TARGETARCH}.tar.gz

COPY --from=docker/buildx-bin /buildx /usr/libexec/docker/cli-plugins/docker-buildx

WORKDIR $GOPATH



