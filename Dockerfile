FROM golang:1.19.4

RUN apt-get update
RUN apt-get install -y aria2 unzip

# must match version in prototool.yml file
ENV PROTOBUF_VERSION=3.13.0
ENV PROTOBUF_CHECKSUM=fbebe5e32db9edbb1bf7988af5fed471d22730104bc6ebd5066c5b4646a0949e49139382cae2605c7abc188ea53f73b044f688c264726009fc68cbbab6a98819

# must match versions in tools/go.mod
ENV GO_PROTO_VALIDATORS_VERSION=0.3.2

RUN mkdir /tmp/protoc
RUN echo https://github.com/protocolbuffers/protobuf/releases/download/v${PROTOBUF_VERSION}/protoc-${PROTOBUF_VERSION}-linux-x86_64.zip
RUN aria2c https://github.com/protocolbuffers/protobuf/releases/download/v${PROTOBUF_VERSION}/protoc-${PROTOBUF_VERSION}-linux-x86_64.zip \
  --checksum=sha-512=${PROTOBUF_CHECKSUM} \
  --dir=/tmp/protoc --out=protoc.zip
RUN unzip /tmp/protoc/protoc.zip -d /tmp/protoc
RUN mv -v /tmp/protoc/include/* /usr/local/include
RUN mv -v /tmp/protoc/bin/* /usr/local/bin
RUN rm -frv /tmp/protoc

RUN mkdir /tmp/go
COPY tools/go.mod tools/go.sum tools/tools.go /tmp/go/
RUN cd /tmp/go && go install -v -mod=readonly \
  github.com/golang/protobuf/protoc-gen-go \
  github.com/mwitkow/go-proto-validators/protoc-gen-govalidators \
  github.com/uber/prototool/cmd/prototool \
  golang.org/x/tools/cmd/goimports
RUN mv -v /go/bin/* /usr/local/bin
RUN mkdir -p /usr/local/include/github.com/mwitkow/go-proto-validators
RUN mv -v /go/pkg/mod/github.com/mwitkow/go-proto-validators@v${GO_PROTO_VALIDATORS_VERSION}/*.proto /usr/local/include/github.com/mwitkow/go-proto-validators
RUN go clean -cache
RUN go clean -modcache
RUN rm -frv /go

ENV PROTOTOOL_PROTOC_BIN_PATH=/usr/local/bin/protoc
ENV PROTOTOOL_PROTOC_WKT_PATH=/usr/local/include

RUN protoc --version
RUN prototool version

WORKDIR /work
LABEL org.opencontainers.image.source https://github.com/percona-platform/dbaas-api
