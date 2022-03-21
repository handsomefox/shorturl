FROM golang:1.18 AS builder
RUN go version
COPY . "/go/src/github.com/handsomefox/shorturl"
WORKDIR "/go/src/github.com/handsomefox/shorturl"
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "-s -w" -o /shorturl
CMD ["/shorturl"]
EXPOSE 8000

FROM scratch
COPY --from=builder /shorturl .

EXPOSE 8000

CMD ["/shorturl"]
