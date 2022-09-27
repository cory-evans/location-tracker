FROM golang:1.19 as builder

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY deviceauth ./deviceauth/
COPY devicelocation ./devicelocation/

COPY main.go ./

RUN go build -o /server main.go

FROM alpine as runner

RUN apk add --no-cache \
    ca-certificates \
    unzip \
    wget \
    zip \
    zlib-dev \
    bash

RUN mkdir -p /pb_data

ADD https://github.com/benbjohnson/litestream/releases/download/v0.3.8/litestream-v0.3.8-linux-amd64-static.tar.gz /tmp/litestream.tar.gz
RUN tar -C /usr/local/bin -xzf /tmp/litestream.tar.gz && rm /tmp/litestream.tar.gz

EXPOSE 8090

COPY ./scripts/run.sh /scripts/run.sh
RUN chmod +x /scripts/run.sh

COPY --from=builder /server /server

CMD [ "/scripts/run.sh" ]