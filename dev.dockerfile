ARG GO_VERSION=1.18
FROM golang:${GO_VERSION}-alpine AS build_base
LABEL stage=build_base
RUN apk update && apk add gcc libc-dev make git --no-cache ca-certificates  && \
    mkdir /user && \
    echo 'nobody:x:65534:65534:nobody:/:' > /user/passwd && \
    echo 'nobody:x:65534:' > /user/group

RUN mkdir -p /go/src/wagerservice
WORKDIR /go/src/wagerservice

ENV GO111MODULE=on
COPY ./go.mod .
COPY ./go.sum .
RUN go mod download

# ================ copy from stage build ===========

FROM build_base AS server_builder
LABEL stage=server_builder

RUN mkdir -p /go/src/wagerservice/cmd
COPY ./cmd/entity-server /go/src/wagerservice/cmd/entity-server

RUN mkdir -p /go/src/wagerservice/internal
COPY ./internal/pkg /go/src/wagerservice/internal/pkg

RUN mkdir -p /go/src/wagerservice/config
COPY ./config /go/src/wagerservice/config

WORKDIR /go/src/wagerservice/cmd/entity-server/

RUN go build  -o /entity-server .

# ================ copy from stage build ===========
FROM alpine:3.8

RUN apk update &&  apk add --no-cache ca-certificates git && \
    mkdir /user && \
    echo 'nobody:x:65534:65534:nobody:/:' > /user/passwd && \
    echo 'nobody:x:65534:' > /user/group

RUN mkdir -p /wagerservice
WORKDIR /wagerservice

RUN mkdir -p /wagerservice/cmd/bin
RUN mkdir -p /wagerservice/config

COPY --from=server_builder /entity-server /wagerservice/cmd/bin/
COPY --from=server_builder /go/src/wagerservice/config/config.toml /wagerservice/config/

RUN chmod -R 777 /wagerservice/cmd/bin

RUN chown -R nobody:nobody /wagerservice
RUN chmod -R 755 /wagerservice

USER nobody:nobody

EXPOSE 8080
ENTRYPOINT ["/wagerservice/cmd/bin/entity-server"]