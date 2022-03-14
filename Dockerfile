FROM golang:1.16-buster as builder

ARG http_proxy
ARG https_proxy
ARG no_proxy
ARG GOPROXY
ARG GOPRIVATE
ARG CI_JOB_TOKEN
ARG CI_COMMIT_SHA

RUN apt-get update \
    && apt-get install ca-certificates -y \
    && update-ca-certificates --fresh \
    && apt-get install -y git

ENV GO111MODULE "on"

WORKDIR /app

COPY go.mod .
COPY go.sum .

COPY . .

RUN GOOS=linux GOARCH=amd64 go build \
    -o main \
    -ldflags "-X main.CommitSHA=$CI_COMMIT_SHA -X main.BuildDatetime=$(date --iso-8601=seconds)" \
    ./cmd/main


FROM buster-slim

ARG http_proxy
ARG https_proxy
ARG no_proxy

WORKDIR /app

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /etc/passwd /etc/passwd

COPY --from=builder /app/main /app/

ENTRYPOINT ["/app/main"]
