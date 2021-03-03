FROM golang:alpine

ENV GO111MODULE=on
ENV GOPROXY=https://goproxy.io,direct

WORKDIR /go/src/pep
COPY . .
RUN go env && go build -o server .

FROM alpine:latest
LABEL MAINTAINER="SliverHorn@sliver_horn@qq.com"

WORKDIR /go/src/pep
COPY --from=0 /go/src/pep/server ./
COPY --from=0 /go/src/pep/config.yaml ./
COPY --from=0 /go/src/pep/resource ./resource

EXPOSE 8888

ENTRYPOINT ./server
