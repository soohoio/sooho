# syntax = docker/dockerfile:1

ARG GO_VERSION="1.19"
ARG RUNNER_IMAGE="alpine:3.16"

FROM golang:${GO_VERSION}-alpine as builder

WORKDIR /opt
RUN apk add --no-cache make git gcc musl-dev openssl-dev linux-headers

COPY go.mod .
COPY go.sum .

RUN --mount=type=cache,target=/root/.cache/go-build \
    --mount=type=cache,target=/root/go/pkg/mod \
    go mod download

# Copy the remaining files
COPY . .

RUN LINK_STATICALLY=true make build

# Add to a distroless container
FROM ${RUNNER_IMAGE}

COPY --from=builder /opt/build/staykingd /usr/local/bin/staykingd
RUN apk add bash vim sudo dasel \
    && addgroup -g 1000 stayking \
    && adduser -S -h /home/stayking -D stayking -u 1000 -G stayking

RUN mkdir -p /etc/sudoers.d \
    && echo '%wheel ALL=(ALL) ALL' > /etc/sudoers.d/wheel \
    && echo "%wheel ALL=(ALL) NOPASSWD: ALL" > /etc/sudoers \
    && adduser stayking wheel

USER 1000
ENV HOME /home/stayking
WORKDIR $HOME

EXPOSE 26657 26656 1317 9090

CMD ["staykingd", "start"]