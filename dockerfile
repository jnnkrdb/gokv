# ---------------------------------------------------------------------------------------------- Golang
# Building the go binary
FROM golang:1.24.2 AS builder
RUN mkdir -p /github.com/jnnkrdb/gokv
WORKDIR /github.com/jnnkrdb/gokv
# copy the code files
COPY gokv/ ./
# set env vars
ENV CGO_ENABLED=0
ENV GOARCH=amd64
ENV GOOS=linux
# START BUILD
RUN go mod download && go build -o /gokv /github.com/jnnkrdb/gokv/main.go
# ---------------------------------------------------------------------------------------------- Final Alpine
FROM alpine:3.22.0
LABEL org.opencontainers.image.source="https://github.com/jnnkrdb/gokv"
LABEL org.opencontainers.image.author="jnnkrdb"
LABEL org.opencontainers.image.description="KeyValue Store in Go."
WORKDIR /
# install neccessary binaries
RUN apk add --no-cache --update openssl
# Copy the Directory Contents
RUN mkdir /opt/gokv
COPY gokv/ /opt/gokv
# create user with home dir
RUN addgroup -S gokv && adduser -S gokv -H -h /opt/gokv -s /bin/sh -G gokv -u 3453
# Copy Binary
COPY --from=builder /gokv /usr/local/bin/gokv
RUN chmod 700 /usr/local/bin/gokv &&\
    chmod 700 -R /opt/gokv &&\
    chown gokv:gokv /usr/local/bin/gokv &&\
    chown gokv:gokv -R /opt/gokv
# change to required user
USER gokv:gokv
# set env vars for the binary
ENV GOKV_HOME="/opt/gokv"
ENV GOKV_BINARY_PATH="/usr/local/bin/gokv"
# set the entrypoints
ENTRYPOINT ["/opt/gokv/entrypoint.sh"]
CMD [ "gokv" ]