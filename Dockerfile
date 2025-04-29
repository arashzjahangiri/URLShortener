# STEP 1 build executable binary
FROM golang:1.24.2 AS builder
WORKDIR /src
COPY . ./
RUN cd cmd/server && CGO_ENABLED=0 go build -o URLShortener && mv URLShortener /usr/bin

# STEP 2 build a small image
FROM alpine:3.20
LABEL maintainer="Mohammad Nasr <arashzjahangiri.dev@gmail.com>"
RUN apk add --no-cache bind-tools busybox-extras
COPY --from=builder /usr/bin/URLShortener /usr/bin/URLShortener
ENV USER=root
ENTRYPOINT ["/usr/bin/URLShortener"]
