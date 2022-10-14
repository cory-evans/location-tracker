FROM --platform=$BUILDPLATFORM golang:1.19 as builder
ARG TARGETARCH
ARG BUILDPLATFORM

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY devicelocation ./devicelocation/

COPY main.go ./

RUN GOOS=linux GOARCH=$TARGETARCH go build -o /server main.go

FROM alpine as runner

ARG LITESTREAM_VERSION=0.3.9
ARG TARGETPLATFORM

RUN apk add --no-cache \
    ca-certificates \
    unzip \
    wget \
    zip \
    zlib-dev \
    bash

RUN mkdir -p /pb_data

RUN wget -O /tmp/litestream.tar.gz https://github.com/benbjohnson/litestream/releases/download/v${LITESTREAM_VERSION}/litestream-v${LITESTREAM_VERSION}-$(echo "$TARGETPLATFORM" | tr '/' '-')-static.tar.gz
RUN tar -C /usr/local/bin -xzf /tmp/litestream.tar.gz && rm /tmp/litestream.tar.gz

EXPOSE 8090

COPY ./scripts/run.sh /scripts/run.sh
RUN chmod +x /scripts/run.sh

COPY --from=builder /server /server

CMD [ "/scripts/run.sh" ]