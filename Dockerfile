FROM golang:1.18-alpine AS builder
RUN apk update
RUN apk add --no-cache git
RUN apk add -U --no-cache ca-certificates && update-ca-certificates
COPY . "/go/src/github.com/handsomefox/shorturl"
WORKDIR "/go/src/github.com/handsomefox/shorturl"
ENV GO111MODULE=on
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "-s -w" -o /shorturl

FROM scratch
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /shorturl .

EXPOSE $PORT
ENTRYPOINT ["/shorturl"]
