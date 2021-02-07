FROM golang:1.15.0

ARG PROTO_VERSION=3.14.0
ARG GRPCWEB_VERSION=1.0.7

WORKDIR /proto

# install unzip
RUN apt-get -qq update
RUN apt-get -qq install -y unzip

# install protoc
RUN curl -sLJO https://github.com/protocolbuffers/protobuf/releases/download/v${PROTO_VERSION}/protoc-${PROTO_VERSION}-linux-x86_64.zip
RUN unzip -qq protoc-${PROTO_VERSION}-linux-x86_64.zip
RUN cp ./bin/protoc /usr/local/bin/protoc
RUN cp -r ./include /usr/local

# install protoc-gen-grpc-web
RUN curl -sLJO https://github.com/grpc/grpc-web/releases/download/${GRPCWEB_VERSION}/protoc-gen-grpc-web-${GRPCWEB_VERSION}-linux-x86_64
RUN cp protoc-gen-grpc-web-${GRPCWEB_VERSION}-linux-x86_64 /usr/local/bin/protoc-gen-grpc-web
RUN chmod +x /usr/local/bin/protoc-gen-grpc-web

RUN go get -u github.com/golang/protobuf/protoc-gen-go
