FROM golang:buster as build

RUN apt-get update && apt-get install -y \
	libwebsockets8 \
	libwebsockets-dev \
	libc-ares2 \
	libc-ares-dev \
	openssl \
	uuid \
	uuid-dev

WORKDIR /work

RUN wget -O mosquitto-2.0.15.tar.gz https://github.com/eclipse/mosquitto/archive/refs/tags/v2.0.15.tar.gz && \
	tar xzvf mosquitto-2.0.15.tar.gz && \
	cp mosquitto-2.0.15/include/* /usr/local/include/ \
	&& rm -Rf mosquitto-2.0.15.tar.gz

RUN wget -O mosquitto-go-auth-2.0.0.tar.gz https://github.com/iegomez/mosquitto-go-auth/archive/refs/tags/2.0.0.tar.gz \
	&& ls -l \
	&& tar xvf *.tar.gz --strip-components=1 \
	&& rm -Rf *.tar.gz

RUN make

FROM ubuntu:latest
ENV USER=mosquitto \
    UID=1000 \
    GID=1000

RUN addgroup --gid "$GID" "$USER" \
    && adduser \
    --disabled-password \
    --gecos "" \
    --home "/mosquitto" \
    --ingroup "$USER" \
    --no-create-home \
    --uid "$UID" \
    "$USER"

RUN apt-get update && \
	apt-get install -y mosquitto

RUN mkdir -p /etc/mosquitto/conf.d
COPY --from=build /work/go-auth.so /etc/mosquitto/conf.d/go-auth.so

WORKDIR /mosquitto
USER mosquitto

ENTRYPOINT [ "mosquitto" ]
CMD ["-c", "/etc/mosquitto/mosquitto.conf" ]