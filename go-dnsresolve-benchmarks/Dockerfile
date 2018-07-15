FROM golang:1.9 as builder
WORKDIR /go/src/github.com/rocketspacer/go-dnsresolve-benchmark
COPY *.go .
RUN CGO_ENABLED=0 GOOS=linux go build -a -o exec .

FROM ubuntu:xenial
RUN apt update && apt install -y ca-certificates
WORKDIR /app/
COPY --from=builder /go/src/github.com/rocketspacer/go-dnsresolve-benchmark/exec .
CMD ["./exec"]